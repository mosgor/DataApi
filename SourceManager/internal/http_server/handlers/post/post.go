package post

import (
	"SourceManager/internal/config"
	"SourceManager/internal/db"
	"SourceManager/internal/logger"
	"SourceManager/internal/repositories"
	"encoding/json"
	"log/slog"
	"net/http"
)

func ServerPost(log *slog.Logger, repository repositories.SourceRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.MustLoad()
		log := logger.SetupLogger(cfg.Env)
		var newData db.Source
		err := json.NewDecoder(r.Body).Decode(&newData)
		if err != nil {
			log.Error("FROM http_server/handlers/post", err.Error(), err)
			return
		}

		repository.Create(r.Context(), &newData)
		// err := json.NewDecoder(r.Body).Decode(&newData)
		// if err != nil {
		// 	log.Error(err.Error())
		// 	return
		// }
	})
}
