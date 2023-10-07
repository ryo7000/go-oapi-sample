package repository

import (
	"context"

	"github.com/ryo7000/go-oapi-sample/domain/model"
)

type ArticleRepository interface {
	Create(c context.Context, article *model.Article) error
	Find(c context.Context, id uint64) (*model.Article, error)
	List(c context.Context) ([]*model.Article, error)
	Update(c context.Context, article *model.Article) (*model.Article, error)
	Delete(c context.Context, id uint64) error
}
