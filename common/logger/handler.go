package logger

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
)

type _CtxKeyType struct{}

var ctxKey = _CtxKeyType{}

func WithRequestID(ctx echo.Context, requestID string) context.Context {
	return context.WithValue(ctx.Request().Context(), ctxKey, requestID)
}

type LoggerHandler struct {
	parent slog.Handler
}

func NewLoggerHandler(parent slog.Handler) *LoggerHandler {
	return &LoggerHandler{parent}
}

func (h *LoggerHandler) Handle(ctx context.Context, record slog.Record) error {
	requestID := ctx.Value(ctxKey)
	if requestID != nil {
		record.Add(slog.String("request-id", requestID.(string)))
	}

	return h.parent.Handle(ctx, record)
}

func (h *LoggerHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.parent.Enabled(ctx, level)
}

func (h *LoggerHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LoggerHandler{parent: h.parent.WithAttrs(attrs)}
}

func (h *LoggerHandler) WithGroup(name string) slog.Handler {
	return &LoggerHandler{parent: h.parent.WithGroup(name)}
}
