package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetTags(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
