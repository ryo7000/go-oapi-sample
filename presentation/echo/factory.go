package echo

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/ryo7000/go-oapi-sample/domain/model"
)

func newUser(req *NewUserRequest) *model.User {
	return &model.User{
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password,
	}
}

func newTag(description string) *model.Tag {
	return &model.Tag{
		Description: description,
	}
}

func newArticle(req *NewArticleRequest) *model.Article {
	article := &model.Article{
		Body:        req.Article.Body,
		Description: req.Article.Description,
	}

	for _, tag := range *req.Article.TagList {
		article.Tags = append(article.Tags, *newTag(tag))
	}

	return article
}

func newUserResponse(u *model.User) (UserResponse, error) {
	token, err := Sign(u.ID)
	if err != nil {
		return UserResponse{}, err
	}

	user := User{
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
	}

	if u.Bio != nil {
		user.Bio = *u.Bio
	}

	if u.Image != nil {
		user.Image = *u.Image
	}

	return UserResponse{
		User: user,
	}, nil
}

func newArticleResponse(userID uint64, a *model.Article) *SingleArticleResponse {
	article := Article{}
	if a.Author.Bio != nil {
		article.Author.Bio = *a.Author.Bio
	}
	article.Author.Following = a.Author.FollowedBy(userID)
	if a.Author.Image != nil {
		article.Author.Image = *a.Author.Image
	}
	article.Author.Username = a.Author.Username

	article.Body = a.Body
	article.CreatedAt = a.CreatedAt
	article.Description = a.Description
	for _, u := range a.Favorites {
		if u.ID == userID {
			article.Favorited = true
		}
	}
	article.FavoritesCount = len(a.Favorites)
	article.Slug = a.Slug
	for _, t := range a.Tags {
		article.TagList = append(article.TagList, t.Description)
	}
	article.Title = a.Title
	article.UpdatedAt = a.UpdatedAt

	return &SingleArticleResponse{article}
}

func newGenericError(ctx echo.Context, err error) GenericError {
	slog.WarnContext(ctx.Request().Context(), "Unexpected error", "err", err)

	genErr := GenericError{}
	genErr.Errors.Body = []string{"Unexpected error"}

	return genErr
}
