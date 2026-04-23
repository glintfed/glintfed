package main

import (
	"glintfed/internal/data/client"
	"glintfed/internal/server"
	"net/http"

	"github.com/mazrean/kessoku"
)

//go:generate go tool kessoku $GOFILE
var _ = kessoku.Inject[*App](
	"newApp",
	kessoku.Provide(client.NewDatabase),
	kessoku.Provide(server.NewAPIServer),
	kessoku.Provide(func(srv *http.Server) *App { return &App{HTTPServer: srv} }),
)
