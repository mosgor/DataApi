package logs

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/structs"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewFindAll(log *slog.Logger, repository repositories.LogsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "logs.handlers.NewFindAll"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		res, err := repository.ReadAll(r.Context())
		if err != nil {
			log.Error("Failed to read logs", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, struct{ Models []structs.LogsModel }{
			Models: res,
		})
	}
}
