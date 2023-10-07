package repository

import (
	"context"

	"github.com/ryo7000/go-oapi-sample/domain/model"
	"github.com/ryo7000/go-oapi-sample/domain/repository"
	"gorm.io/gorm"
)

type articleRepository struct {
	Conn *gorm.DB
}

func NewArticleRepository(Conn *gorm.DB) repository.ArticleRepository {
	return &articleRepository{Conn}
}

func (r *articleRepository) Create(c context.Context, article *model.Article) error {
	return r.Conn.WithContext(c).Create(article).Error
}

func (r *articleRepository) Find(c context.Context, id uint64) (*model.Article, error) {
	a := &model.Article{ID: id}
	err := r.Conn.WithContext(c).First(a).Error
	return a, err
}

func (r *articleRepository) List(c context.Context) ([]*model.Article, error) {
	var articles []*model.Article
	err := r.Conn.WithContext(c).Find(&articles).Error
	return articles, err
}

func (r *articleRepository) Update(c context.Context, article *model.Article) (*model.Article, error) {
	err := r.Conn.WithContext(c).Model(article).Updates(article).Error
	return article, err
}

func (r *articleRepository) Delete(c context.Context, id uint64) error {
	a := &model.Article{ID: id}
	err := r.Conn.WithContext(c).Delete(a).Error
	return err
}
