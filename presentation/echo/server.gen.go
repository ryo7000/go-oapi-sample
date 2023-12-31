// Package echo provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package echo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get recent articles globally
	// (GET /articles)
	GetArticles(ctx echo.Context, params GetArticlesParams) error
	// Create an article
	// (POST /articles)
	CreateArticle(ctx echo.Context) error
	// Get recent articles from users you follow
	// (GET /articles/feed)
	GetArticlesFeed(ctx echo.Context, params GetArticlesFeedParams) error
	// Delete an article
	// (DELETE /articles/{slug})
	DeleteArticle(ctx echo.Context, slug string) error
	// Get an article
	// (GET /articles/{slug})
	GetArticle(ctx echo.Context, slug string) error
	// Update an article
	// (PUT /articles/{slug})
	UpdateArticle(ctx echo.Context, slug string) error
	// Get comments for an article
	// (GET /articles/{slug}/comments)
	GetArticleComments(ctx echo.Context, slug string) error
	// Create a comment for an article
	// (POST /articles/{slug}/comments)
	CreateArticleComment(ctx echo.Context, slug string) error
	// Delete a comment for an article
	// (DELETE /articles/{slug}/comments/{id})
	DeleteArticleComment(ctx echo.Context, slug string, id int) error
	// Unfavorite an article
	// (DELETE /articles/{slug}/favorite)
	DeleteArticleFavorite(ctx echo.Context, slug string) error
	// Favorite an article
	// (POST /articles/{slug}/favorite)
	CreateArticleFavorite(ctx echo.Context, slug string) error
	// Get a profile
	// (GET /profiles/{username})
	GetProfileByUsername(ctx echo.Context, username string) error
	// Unfollow a user
	// (DELETE /profiles/{username}/follow)
	UnfollowUserByUsername(ctx echo.Context, username string) error
	// Follow a user
	// (POST /profiles/{username}/follow)
	FollowUserByUsername(ctx echo.Context, username string) error
	// Get tags
	// (GET /tags)
	GetTags(ctx echo.Context) error
	// Get current user
	// (GET /user)
	GetCurrentUser(ctx echo.Context) error
	// Update current user
	// (PUT /user)
	UpdateCurrentUser(ctx echo.Context) error

	// (POST /users)
	CreateUser(ctx echo.Context) error
	// Existing user login
	// (POST /users/login)
	Login(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArticles converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticles(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetArticlesParams
	// ------------- Optional query parameter "tag" -------------

	err = runtime.BindQueryParameter("form", true, false, "tag", ctx.QueryParams(), &params.Tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tag: %s", err))
	}

	// ------------- Optional query parameter "author" -------------

	err = runtime.BindQueryParameter("form", true, false, "author", ctx.QueryParams(), &params.Author)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter author: %s", err))
	}

	// ------------- Optional query parameter "favorited" -------------

	err = runtime.BindQueryParameter("form", true, false, "favorited", ctx.QueryParams(), &params.Favorited)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter favorited: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticles(ctx, params)
	return err
}

// CreateArticle converts echo context to params.
func (w *ServerInterfaceWrapper) CreateArticle(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateArticle(ctx)
	return err
}

// GetArticlesFeed converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticlesFeed(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetArticlesFeedParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticlesFeed(ctx, params)
	return err
}

// DeleteArticle converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteArticle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteArticle(ctx, slug)
	return err
}

// GetArticle converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticle(ctx, slug)
	return err
}

// UpdateArticle converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateArticle(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateArticle(ctx, slug)
	return err
}

// GetArticleComments converts echo context to params.
func (w *ServerInterfaceWrapper) GetArticleComments(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetArticleComments(ctx, slug)
	return err
}

// CreateArticleComment converts echo context to params.
func (w *ServerInterfaceWrapper) CreateArticleComment(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateArticleComment(ctx, slug)
	return err
}

// DeleteArticleComment converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteArticleComment(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteArticleComment(ctx, slug, id)
	return err
}

// DeleteArticleFavorite converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteArticleFavorite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteArticleFavorite(ctx, slug)
	return err
}

// CreateArticleFavorite converts echo context to params.
func (w *ServerInterfaceWrapper) CreateArticleFavorite(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithLocation("simple", false, "slug", runtime.ParamLocationPath, ctx.Param("slug"), &slug)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter slug: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateArticleFavorite(ctx, slug)
	return err
}

// GetProfileByUsername converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfileByUsername(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProfileByUsername(ctx, username)
	return err
}

// UnfollowUserByUsername converts echo context to params.
func (w *ServerInterfaceWrapper) UnfollowUserByUsername(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UnfollowUserByUsername(ctx, username)
	return err
}

// FollowUserByUsername converts echo context to params.
func (w *ServerInterfaceWrapper) FollowUserByUsername(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FollowUserByUsername(ctx, username)
	return err
}

// GetTags converts echo context to params.
func (w *ServerInterfaceWrapper) GetTags(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTags(ctx)
	return err
}

// GetCurrentUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetCurrentUser(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCurrentUser(ctx)
	return err
}

// UpdateCurrentUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCurrentUser(ctx echo.Context) error {
	var err error

	ctx.Set(TokenScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateCurrentUser(ctx)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/articles", wrapper.GetArticles)
	router.POST(baseURL+"/articles", wrapper.CreateArticle)
	router.GET(baseURL+"/articles/feed", wrapper.GetArticlesFeed)
	router.DELETE(baseURL+"/articles/:slug", wrapper.DeleteArticle)
	router.GET(baseURL+"/articles/:slug", wrapper.GetArticle)
	router.PUT(baseURL+"/articles/:slug", wrapper.UpdateArticle)
	router.GET(baseURL+"/articles/:slug/comments", wrapper.GetArticleComments)
	router.POST(baseURL+"/articles/:slug/comments", wrapper.CreateArticleComment)
	router.DELETE(baseURL+"/articles/:slug/comments/:id", wrapper.DeleteArticleComment)
	router.DELETE(baseURL+"/articles/:slug/favorite", wrapper.DeleteArticleFavorite)
	router.POST(baseURL+"/articles/:slug/favorite", wrapper.CreateArticleFavorite)
	router.GET(baseURL+"/profiles/:username", wrapper.GetProfileByUsername)
	router.DELETE(baseURL+"/profiles/:username/follow", wrapper.UnfollowUserByUsername)
	router.POST(baseURL+"/profiles/:username/follow", wrapper.FollowUserByUsername)
	router.GET(baseURL+"/tags", wrapper.GetTags)
	router.GET(baseURL+"/user", wrapper.GetCurrentUser)
	router.PUT(baseURL+"/user", wrapper.UpdateCurrentUser)
	router.POST(baseURL+"/users", wrapper.CreateUser)
	router.POST(baseURL+"/users/login", wrapper.Login)

}
