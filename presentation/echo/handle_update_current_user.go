package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) UpdateCurrentUser(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
