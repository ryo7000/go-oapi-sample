package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ryo7000/go-oapi-sample/domain/model"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	_db, purge := CreateDB()
	db = _db

	code := m.Run()

	if err := purge(); err != nil {
		log.Fatalf("Could not purge: %s", err)
	}

	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	repo := NewUserRepository(db)

	user := &model.User{Username: "foobar", Email: "foobar@example.com", Password: "abc"}
	hashed, err := model.HashPassword(user.Password)
	if err != nil {
		t.Fatal("Password generate err", err)
		return
	}

	user.Password = hashed

	err = repo.Create(context.Background(), user)
	if err != nil {
		t.Fatal("Could not create user", err)
		return
	}
	if user.ID == 0 {
		t.Error("ID not created")
	}
}
