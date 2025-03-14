package logger

import (
	"log/slog"
	"os"
)

type AppLogger struct {
	Logger *slog.Logger
}

// A NewAppLogger létrehozza az AppLogger új példányát.
// Az AppLogger struktúra tartalmaz egy slog.Logger példányt a naplózáshoz
// Parameters:
// -
// Returns:
// - *AppLogger: Aplogger új példány mutatója
func NewAppLogger() *AppLogger {
	return &AppLogger{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})),
	}
}
