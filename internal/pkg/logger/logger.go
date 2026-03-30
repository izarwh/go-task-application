package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	Level       string
	Environment string
}

func Init(cfg Config) {
	var level slog.Level
	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	}

	if cfg.Environment == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	handler = &ContextHandler{handler}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
