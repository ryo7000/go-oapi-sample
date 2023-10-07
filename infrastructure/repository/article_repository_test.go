package repository

import (
	"context"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ryo7000/go-oapi-sample/domain/model"
)

func TestCreateArticle(t *testing.T) {
	repo := NewArticleRepository(db)

	article := &model.Article{
		Title: "title",
	}

	err := repo.Create(context.Background(), article)
	if err != nil {
		t.Fatal("Could not create article", err)
		return
	}
	if article.ID == 0 {
		t.Error("ID not created")
	}
}
