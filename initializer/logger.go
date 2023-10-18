package initializer

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"strings"
	"thinkprinter/models"
)

func initLogger() {
	w := os.Stderr

	opts := &tint.Options{
		AddSource:  toLogLevel(models.C.Core.LogLevel) == slog.LevelDebug,
		Level:      toLogLevel(models.C.Core.LogLevel),
		TimeFormat: "2006-01-02 15:04:05",
	}

	logger := slog.New(tint.NewHandler(w, opts))

	slog.SetDefault(logger)
}

func toLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
