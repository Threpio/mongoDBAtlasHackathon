package main

import (
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend/user"
	"net/http"
)

func NewHandler(logger logger.logger) http.Handler {
	// TODO: Sort out Logger

	r := chi.NewRouter()

	// TODO: Add middleware for auth and such

	// User Functions

	// Logging Functions

	// Ingest

	// Eggres

	// Administion Functions


	return r
}
