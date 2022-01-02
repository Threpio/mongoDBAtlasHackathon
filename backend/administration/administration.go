package administration

import (
	"github.com/go-chi/chi"
	"github.com/threpio/mongoDBAtlasHackathon/backend/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend/logger"
)

type Controller struct {
	DB     db.DB
	Router func(r chi.Router)
	logger logger.Logger
}

func NewController(db db.DB, logger logger.Logger) (*Controller, error) {
	controller := &Controller{
		DB:     db,
		logger: logger,
	}
	controller.Router = controller.administrationRouter()
	return controller, nil
}
