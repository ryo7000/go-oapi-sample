// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"github.com/ryo7000/go-oapi-sample/common/logger"
	"github.com/ryo7000/go-oapi-sample/infrastructure"
	"github.com/ryo7000/go-oapi-sample/infrastructure/repository"
	"github.com/ryo7000/go-oapi-sample/presentation/echo"
	"github.com/ryo7000/go-oapi-sample/usecase"
)

// Injectors from wire.go:

func CreateServer() echo.Server {
	slogLogger := logger.NewLogger()
	gormDB := db.NewDB(slogLogger)
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	articleRepository := repository.NewArticleRepository(gormDB)
	articleUseCase := usecase.NewArticleUseCase(articleRepository)
	server := echo.NewServer(slogLogger, userUseCase, articleUseCase)
	return server
}
