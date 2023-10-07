package usecase

import (
	"context"

	"github.com/ryo7000/go-oapi-sample/domain/model"
	"github.com/ryo7000/go-oapi-sample/domain/repository"
)

type ArticleUseCase interface {
	Create(c context.Context, newArticle *model.Article) error
	List(c context.Context, offset int, limit int) ([]*model.Article, error)
}

type articleUseCase struct {
	articleRepo repository.ArticleRepository
}

func NewArticleUseCase(articleRepo repository.ArticleRepository) ArticleUseCase {
	return &articleUseCase{
		articleRepo: articleRepo,
	}
}

func (repo *articleUseCase) Create(c context.Context, newArticle *model.Article) error {
	return repo.articleRepo.Create(c, newArticle)
}

func (repo *articleUseCase) List(c context.Context, offset int, limit int) ([]*model.Article, error) {
	return repo.articleRepo.List(c)
}
