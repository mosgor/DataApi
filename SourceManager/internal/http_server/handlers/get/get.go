package get

import (
	"SourceManager/internal/config"
	"SourceManager/internal/db"
	"SourceManager/internal/logger"
	"SourceManager/internal/repositories"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	db.MongoData
	db.PostgresData
}

var data []Response

func ServerGetOne(log *slog.Logger, repository repositories.SourceRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.MustLoad()
		log := logger.SetupLogger(cfg.Env)
		sourceId, err := strconv.Atoi(chi.URLParam(r, "sourceId"))
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		read, err := repository.ReadOne(r.Context(), sourceId)

		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(read)

		log.Info(fmt.Sprintf("Data sent successfully:%v", data))
	})
}

func ServerGetMultiple(log *slog.Logger, repository repositories.SourceRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.MustLoad()
		log := logger.SetupLogger(cfg.Env)

		read, err := repository.ReadAll(r.Context())

		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(read)

		log.Info(fmt.Sprintf("Data sent successfully:%v", data))
	})
}
