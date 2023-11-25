package log

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

func SetDefault() {
	slog.SetDefault(New())
}

func New() *slog.Logger {
	zapL := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		zapL = zap.Must(zap.NewDevelopment())
	}

	options := &zapslog.HandlerOptions{
		AddSource: false,
	}

	return slog.New(zapslog.NewHandler(zapL.Core(), options))
}
