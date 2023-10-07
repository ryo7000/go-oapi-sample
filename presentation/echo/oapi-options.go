package echo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	oapiMiddleware "github.com/oapi-codegen/echo-middleware"
)

func OapiRequestErrorHandler(c echo.Context, err *echo.HTTPError) error {
	if _, ok := err.Internal.(*openapi3filter.SecurityRequirementsError); ok {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusUnprocessableEntity, newGenericError(c, err))
}

func Authenticator(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "Token" {
		return fmt.Errorf("security scheme %s != 'BearerAuth'", input.SecuritySchemeName)
	}

	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	id, err := Verify(jws)
	if err != nil {
		return err
	}

	ec := oapiMiddleware.GetEchoContext(ctx)
	ec.Set(ctxKey, id)

	return nil
}

func NewOapiRequestValidatorOptions() *oapiMiddleware.Options {
	return &oapiMiddleware.Options{
		ErrorHandler: OapiRequestErrorHandler,
		Options: openapi3filter.Options{
			AuthenticationFunc: Authenticator,
		},
	}
}
