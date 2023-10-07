package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) DeleteArticleComment(ctx echo.Context, slug string, id int) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
