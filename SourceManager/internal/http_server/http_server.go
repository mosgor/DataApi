package http_server

import (
	"fmt"
	"net/http"
	"time"

	"SourceManager/internal/config"
	"SourceManager/internal/http_server/handlers/delete"
	"SourceManager/internal/http_server/handlers/get"
	"SourceManager/internal/http_server/handlers/post"
	"SourceManager/internal/http_server/handlers/put"
	mwLogger "SourceManager/internal/http_server/middleware/logger"

	logs "SourceManager/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateServer() (server *http.Server) {
	cfg := config.MustLoad()
	log := logs.SetupLogger("local")
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Get("/", get.ServerGet())
	router.Post("/", post.ServerPost())
	router.Delete("/", delete.ServerDelete())
	router.Put("/", put.ServerPut())

	return &http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.HTTP.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
		Handler:      router,
	}

}
