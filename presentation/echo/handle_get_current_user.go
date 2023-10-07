package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetCurrentUser(ctx echo.Context) error {
	id, err := GetUserId(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
	}

	user, err := s.userUseCase.FindById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
	}

	res, err := newUserResponse(user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
	}

	return ctx.JSON(http.StatusOK, res)
}
