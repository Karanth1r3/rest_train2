package slg

import (
	"log/slog"
	"os"
)

// Wrapping errors to logger
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func LogErrorFatal(l *slog.Logger, message string, err error) {
	l.Error(message, Err(err))
	os.Exit(1)
}
