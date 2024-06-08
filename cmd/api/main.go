package main

import (
	"log/slog"
	"os"

	"github.com/evertonbzr/api_todo/internal/api"
	"github.com/evertonbzr/api_todo/internal/config"
)

func main() {
	config.Load(os.Getenv("ENVIRONMENT"))

	slog.Info("Starting API", "environment", config.ENVIRONMENT, "port", config.PORT)

	apiConf := &api.APIConfig{
		Port: config.PORT,
	}

	api.Run(apiConf)
}
