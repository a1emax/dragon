package tools

import (
	"log/slog"
	"os"

	"github.com/a1emax/youngine/x/scope"
)

var Logger *slog.Logger

func initLogger(lc scope.Lifecycle) {
	Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}
