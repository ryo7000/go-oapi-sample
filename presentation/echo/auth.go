package echo

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const ctxKey = "user_id"

func GetUserId(ec echo.Context) (uint64, error) {
	v := ec.Get(ctxKey)
	id, ok := v.(uint64)
	if !ok {
		return 0, errors.New("unknown user id")
	}

	return id, nil
}

func GetJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")
	if authHdr == "" {
		return "", errors.New("authorization header is missing")
	}

	prefix := "Token "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", errors.New("authorization header is malformed")
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}
