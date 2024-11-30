package http_server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"SourceManager/internal/config"
	"SourceManager/internal/db/sources"
	"SourceManager/internal/http_server/handlers/delete"
	"SourceManager/internal/http_server/handlers/get"
	"SourceManager/internal/http_server/handlers/post"
	"SourceManager/internal/http_server/handlers/put"
	mwLogger "SourceManager/internal/http_server/middleware/logger"

	logs "SourceManager/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateServer() (server *http.Server) {
	cfg := config.MustLoad()
	log := logs.SetupLogger(cfg.Env)
	router := chi.NewRouter()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgresql://admin:%v@localhost:5438/DataApi", cfg.DatabasePass))
	if err != nil {
		log.Error("unable to connect to postgres")
		return
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Error("unable to connect to mongo")
		return
	}

	repo := sources.NewRepository(pool, client, log)

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/source", get.ServerGetMultiple(log, repo))
	router.Get("/source/{sourceId}", get.ServerGetOne(log, repo))

	router.Delete("/source/{sourceID}", delete.ServerDelete())

	router.Put("/source", put.ServerPut())
	router.Post("/source", post.ServerPost())

	return &http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.HTTP.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
		Handler:      router,
	}

}
