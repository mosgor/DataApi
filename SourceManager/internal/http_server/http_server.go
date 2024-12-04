package http_server

import (
	"net/http"
	"time"

	"SourceManager/internal/config"
	"SourceManager/internal/http_server/handlers/get"
	"SourceManager/internal/http_server/handlers/post"
	mwLogger "SourceManager/internal/http_server/middleware/logger"
	"SourceManager/internal/repositories"

	logs "SourceManager/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateServer(repo repositories.SourceRepository) (server *http.Server) {
	cfg := config.MustLoad()
	log := logs.SetupLogger(cfg.Env)
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/source", get.ServerGetMultiple(log, repo))
	router.Get("/source/{sourceId}", get.ServerGetOne(log, repo))

	router.Post("/source", post.ServerPost(log, repo))

	return &http.Server{
		Addr:         cfg.HTTP.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
		Handler:      router,
	}

}
