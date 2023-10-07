package echo

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
	slogecho "github.com/samber/slog-echo"

	"github.com/ryo7000/go-oapi-sample/common/logger"
	"github.com/ryo7000/go-oapi-sample/usecase"
)

type Server struct {
	echo           *echo.Echo
	userUseCase    usecase.UserUseCase
	articleUseCase usecase.ArticleUseCase
}

func NewServer(slogger *slog.Logger, user usecase.UserUseCase, article usecase.ArticleUseCase) Server {
	s := Server{
		echo:           echo.New(),
		userUseCase:    user,
		articleUseCase: article,
	}

	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}

	// Disable verification of servers.url in openapi.yml
	swagger.Servers = nil

	s.echo.Use(slogecho.New(slogger))
	s.echo.Use(oapiMiddleware.OapiRequestValidatorWithOptions(swagger, NewOapiRequestValidatorOptions()))
	s.echo.Use(middleware.RequestIDWithConfig(
		middleware.RequestIDConfig{
			RequestIDHandler: func(c echo.Context, requestID string) {
				ctx := logger.WithRequestID(c, requestID)
				c.SetRequest(c.Request().WithContext(ctx))
			},
		},
	))
	s.echo.Use(middleware.Recover())

	RegisterHandlers(s.echo, &s)

	return s
}

func (s *Server) Start() {
	s.echo.Logger.Fatal(s.echo.Start("0.0.0.0:8080"))
}
