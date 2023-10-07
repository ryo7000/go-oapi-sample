package echo

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateUser(ctx echo.Context) error {
	var req NewUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	user := newUser(&req)
	err := s.userUseCase.Create(ctx.Request().Context(), user)
	if err != nil {
		genErr := GenericError{}

		if errs, ok := err.(validation.Errors); ok {
			for k, err := range errs {
				genErr.Errors.Body = append(genErr.Errors.Body, k+":"+err.Error())
			}
			return ctx.JSON(http.StatusUnprocessableEntity, &genErr)
		} else {
			return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
		}
	}

	res, err := newUserResponse(user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
	}

	return ctx.JSON(http.StatusCreated, res)
}
