package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetProfileByUsername(ctx echo.Context, slug string) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
