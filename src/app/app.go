package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/set0xc3/crmGO/src/config"
	"github.com/set0xc3/crmGO/src/db"
	"github.com/set0xc3/crmGO/src/routes"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()
	srv := routes.NewServer()

	db.Init()
	defer db.DeInit()

	httpServer := &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: srv,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		httpServer.Shutdown(ctx)
	}()

	slog.Info("starting server", slog.String("addr", cfg.ServerAddr))

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
