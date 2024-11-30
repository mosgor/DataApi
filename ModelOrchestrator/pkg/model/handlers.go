package model

import (
	"ModelOrchestrator/pkg/internal/repositories"
	"ModelOrchestrator/pkg/internal/structs"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func NewFindAll(log *slog.Logger, repository repositories.ModelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func NewFindOne(log *slog.Logger, repository repositories.ModelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func NewCreate(log *slog.Logger, repository repositories.ModelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "client.handlers.NewAdd"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req structs.Resp
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Failed to parse request body", err)
			return
		}
		err = repository.Create(r.Context(), &req)
		if err != nil {
			log.Error("Failed to insert model", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, req)
	}
}
