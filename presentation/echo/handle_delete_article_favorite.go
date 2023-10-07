package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) DeleteArticleFavorite(ctx echo.Context, slug string) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
