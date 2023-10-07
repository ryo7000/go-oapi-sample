package main

import (
	"log/slog"

	"github.com/joho/godotenv"

	"github.com/ryo7000/go-oapi-sample/registry"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("Could not load .env file")
	}

	server := registry.CreateServer()
	server.Start()
}
