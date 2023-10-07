package echo

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Login(ctx echo.Context) error {
	var req LoginUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	slog.InfoContext(ctx.Request().Context(), "login", "username", req.User.Email)

	user, err := s.userUseCase.Login(ctx.Request().Context(), req.User.Email, req.User.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
	}

	res, err := newUserResponse(user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
	}

	return ctx.JSON(http.StatusOK, res)
}
