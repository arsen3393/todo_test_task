package logger

import (
	"log/slog"
	"os"
)

func MustInitLogger() *slog.Logger {

	textHandler := slog.NewTextHandler(os.Stdout, nil)

	logger := slog.New(textHandler)

	return logger
}
