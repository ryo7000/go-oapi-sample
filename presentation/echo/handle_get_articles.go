package echo

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ryo7000/go-oapi-sample/domain/model"
)

func toArticle(articles []*model.Article) []Article {
	result := make([]Article, len(articles))

	for _, v := range articles {
		result = append(result, Article{
			Body: v.Body,
		})
	}

	return result
}

func (s *Server) GetArticles(ctx echo.Context, params GetArticlesParams) error {
	slog.InfoContext(ctx.Request().Context(), "Test info log", "username", "foobar")
	articles, err := s.articleUseCase.List(ctx.Request().Context(), 0, 0)
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}

	resArticles := toArticle(articles)
	return ctx.JSON(http.StatusOK, MultipleArticlesResponse{Articles: resArticles, ArticlesCount: len(resArticles)})
}
