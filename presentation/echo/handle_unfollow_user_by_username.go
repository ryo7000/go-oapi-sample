package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) UnfollowUserByUsername(ctx echo.Context, username string) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
