package server

import (
	"glintfed/internal/data"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/go-chi/chi/v5"
)

func NewAPIServer(cfg *data.Config) *http.Server {
	mux := chi.NewRouter()

	return &http.Server{
		Addr:    cfg.Server.API.Addr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
}
