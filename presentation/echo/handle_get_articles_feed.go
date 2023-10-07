package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetArticlesFeed(ctx echo.Context, params GetArticlesFeedParams) error {
	return ctx.NoContent(http.StatusNotImplemented)
}
