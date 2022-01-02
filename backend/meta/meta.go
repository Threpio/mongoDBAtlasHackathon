package meta

import (
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend/logger"
)

type Controller struct {
	logger logger.Logger
	Router func(r chi.Router)
}

func NewController(db db.DB, logger logger.Logger) (*Controller, error) {
	controller := &Controller{
		logger: logger,
	}
	controller.Router = controller.metaRouter()
	return controller, nil
}
