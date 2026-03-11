package logger

import (
	"log/slog"
	"os"

	"github.com/CyberDruga/twitter-bot/src/args"
)

func init() {
	level := slog.LevelInfo

	if *args.Debug {
		level = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		AddSource: *args.Source,
		Level:     level,
	}

	handler := slog.NewTextHandler(os.Stderr, opts)

	slog.SetDefault(slog.New(handler))

}
