package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"glintfed/internal/web"
)

const (
	defaultAddr         = ":8080"
	shutdownGracePeriod = 10 * time.Second
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	server := &http.Server{
		Addr:              listenAddr(),
		Handler:           web.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)

	go func() {
		logger.Printf("starting glintfed on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownGracePeriod)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Fatalf("shutdown server: %v", err)
		}
	case err := <-errCh:
		if err != nil {
			logger.Fatalf("serve http: %v", err)
		}
	}
}

func listenAddr() string {
	if addr := os.Getenv("PORT"); addr != "" {
		if strings.HasPrefix(addr, ":") {
			return addr
		}

		return ":" + addr
	}

	return defaultAddr
}
