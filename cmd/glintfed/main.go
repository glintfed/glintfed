package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"time"

	"glintfed/internal/data"
	"glintfed/internal/lib/liblogs"
)

// Name is the name of the application.
var Name string

// Version is the version of the application.
var Version string

var (
	flagCfgPath         string
	shutdownGracePeriod = 10 * time.Second
)

func init() {
	flag.StringVar(&flagCfgPath, "config", "configs/", "config dir path")
}

func main() {
	flag.Parse()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	cfg, err := data.NewConfig(Name, Version, flagCfgPath)
	if err != nil {
		slog.Error("failed to load config", liblogs.ErrAttr(err))
		os.Exit(1)
	}

	app, err := newApp(cfg)
	if err != nil {
		slog.Error("failed to init application", liblogs.ErrAttr(err))
		return
	}
	if err := app.Run(context.Background()); err != nil {
		slog.Error("failed to run application", liblogs.ErrAttr(err))
		return
	}
}
