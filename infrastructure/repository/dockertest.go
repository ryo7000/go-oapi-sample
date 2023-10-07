package repository

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rubenv/sql-migrate"
)

var (
	dbUser     = "conduit_user"
	dbName     = "conduit_test"
	dbPassword = "password"
	dsn        = "postgres://%s:%s@%s/%s?sslmode=disable"
)

func MigrateDb(db *gorm.DB) error {
	pwd, _ := os.Getwd()
	migrations := &migrate.FileMigrationSource{
		Dir: pwd + "/../../migrations",
	}

	sqlDb, err := db.DB()
	if err != nil {
		return err
	}

	num, err := migrate.Exec(sqlDb, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations!", num)

	return nil
}

func CreateDB() (*gorm.DB, func() error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	options := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.4-alpine",
		Env: []string{
			"POSTGRES_USER=" + dbUser,
			"POSTGRES_PASSWORD=" + dbPassword,
			"POSTGRES_DB=" + dbName,
		},
	}

	resource, err := pool.RunWithOptions(options,
		func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	log.Printf("Postgres running on port %s", hostAndPort)

	var db *gorm.DB
	if err := pool.Retry(func() error {
		var err error
		db, err = gorm.Open(postgres.Open(fmt.Sprintf(dsn, dbUser, dbPassword, hostAndPort, dbName)))
		if err != nil {
			return err
		}

		sqlDb, err := db.DB()
		if err != nil {
			return err
		}

		return sqlDb.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	purge := func() error {
		return pool.Purge(resource)
	}

	err = MigrateDb(db)
	if err != nil {
		log.Fatalf("Could not migrate: %s", err)
	}

	return db, purge
}
