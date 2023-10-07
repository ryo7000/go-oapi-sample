//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/ryo7000/go-oapi-sample/common/logger"
	"github.com/ryo7000/go-oapi-sample/infrastructure"
	"github.com/ryo7000/go-oapi-sample/infrastructure/repository"
	"github.com/ryo7000/go-oapi-sample/presentation/echo"
	"github.com/ryo7000/go-oapi-sample/usecase"

	"github.com/google/wire"
)

func CreateServer() echo.Server {
	wire.Build(
		logger.NewLogger,
		db.NewDB,
		echo.NewServer,
		usecase.NewArticleUseCase,
		usecase.NewUserUseCase,
		repository.NewArticleRepository,
		repository.NewUserRepository,
	)
	return echo.Server{}
}
