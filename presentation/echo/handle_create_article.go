package echo

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateArticle(ctx echo.Context) error {
	var req NewArticleRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	id, err := GetUserId(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, newGenericError(ctx, err))
	}

	_, err = s.userUseCase.FindById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.NoContent(http.StatusUnauthorized)
	}

	article := newArticle(&req)
	article.AuthorID = id
	err = s.articleUseCase.Create(ctx.Request().Context(), article)
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

	return ctx.JSON(http.StatusCreated, newArticleResponse(id, article))
}
