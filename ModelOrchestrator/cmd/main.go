package main

import (
	"ModelOrchestrator/pkg/config"
	"ModelOrchestrator/pkg/gRPC"
	"ModelOrchestrator/pkg/logs"
	"ModelOrchestrator/pkg/mapping"
	"ModelOrchestrator/pkg/model"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"net"
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
	// postgresql://admin:%v@localhost:5438/DataApi for local use
	// postgresql://admin:%v@postgres:5432/DataApi for container deploy
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://admin:%v@postgres:5432/DataApi", cfg.DatabasePass))
	if err != nil {
		log.Error("unable to connect to postgres")
		return
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// mongodb://localhost:27017 for local use
	// mongodb://mongo:27017 for container deploy
	opts := options.Client().ApplyURI("mongodb://mongo:27017").SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Error("unable to connect to mongo")
		return
	}

	modelRepo := model.NewRepository(pool, client, log)
	mappingRepo := mapping.NewRepository(client, log)
	logsRepo := logs.NewRepository(pool, log)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Timeout(cfg.Http.Timeout))

	router.Get("/model", model.NewFindAll(log, modelRepo))
	router.Get("/model/{modelId}", model.NewFindOne(log, modelRepo))
	router.Post("/model", model.NewCreate(log, modelRepo))

	router.Get("/mapping", mapping.NewFindAll(log, mappingRepo))
	router.Get("/mapping/{mappingId}", mapping.NewFindOne(log, mappingRepo))
	router.Post("/mapping", mapping.NewCreate(log, mappingRepo))

	router.Get("/logs", logs.NewFindAll(log, logsRepo))

	netClient := &http.Client{Timeout: cfg.Http.Timeout}
	grpc := gRPC.NewRpcServer(logsRepo, modelRepo, netClient, cfg.Grpc.Timeout, log)
	lis, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		log.Error("Unable to listen to " + cfg.Grpc.Address)
		return
	}
	go func() {
		log.Info("Starting serving gRCP at " + cfg.Grpc.Address)
		err = grpc.Serve(lis)
		if err != nil {
			log.Error("Unable to start gRPC server")
			return
		}
	}()

	srv := &http.Server{
		Addr:         cfg.Http.Address,
		ReadTimeout:  cfg.Http.Timeout,
		WriteTimeout: cfg.Http.Timeout,
		IdleTimeout:  cfg.Http.Timeout,
		Handler:      router,
	}

	log.Info("Starting listening on " + srv.Addr)
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
