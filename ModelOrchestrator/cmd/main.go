package main

import (
	"ModelOrchestrator/pkg/config"
	"ModelOrchestrator/pkg/model"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://admin:%v@localhost:5438/DataApi", cfg.DatabasePass))
	if err != nil {
		log.Error("unable to connect to postgres")
		return
	}

	modelRepo := model.NewRepository(pool, log)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(cfg.Http.Timeout))

	router.Get("/model", model.NewFindAll(log, modelRepo))
	router.Get("/model/{modelId}", model.NewFindOne(log, modelRepo))

	srv := &http.Server{
		Addr:         cfg.Http.Address,
		ReadTimeout:  cfg.Http.Timeout,
		WriteTimeout: cfg.Http.Timeout,
		IdleTimeout:  cfg.Http.Timeout,
		Handler:      router,
	}

	log.Info("Starting listening on ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("can't open server", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
