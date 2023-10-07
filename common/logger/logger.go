package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func NewLogger() *slog.Logger {
	level := new(slog.Level)

	if levelErr := level.UnmarshalText([]byte(os.Getenv("LOG_LEVEL"))); levelErr != nil {
		panic(levelErr)
	}

	logger := slog.New(NewLoggerHandler(tint.NewHandler(os.Stdout, nil)))

	slog.SetDefault(logger)
	return logger
}
