package mapping

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/structs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewFindAll(log *slog.Logger, repository repositories.MappingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "mapping.handlers.NewFindAll"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		res, err := repository.ReadAll(r.Context())
		if err != nil {
			log.Error("Failed to read mappings", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, struct{ Mappings []structs.MappingModel }{
			Mappings: res,
		})
	}
}

func NewFindOne(log *slog.Logger, repository repositories.MappingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "mappings.handlers.NewFindOne"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		mapId := chi.URLParam(r, "mappingId")
		mapping, err := repository.ReadOne(r.Context(), mapId)
		if err != nil {
			log.Error("Failed to get mapping", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, mapping)
	}
}

func NewCreate(log *slog.Logger, repository repositories.MappingRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "mappings.handlers.NewAdd"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req structs.MappingModel
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Failed to parse request body in mapping create", err)
			return
		}
		err = repository.Create(r.Context(), &req)
		if err != nil {
			log.Error("Failed to insert mapping model", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, req)
	}
}
