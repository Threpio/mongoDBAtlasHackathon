package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/ingest"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/logger"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/meta"
)

type App struct {
	Ctx context.Context
	shutdownCallbacks []func()

	IngestController *ingest.Controller
	MetaController   *meta.Controller
}

func NewApp(db db.DB) *App {

	log := logger.New("DEBUG")

	//TODO: Remove this once Debugging
	fmt.Println("Initialising NewApp")
	ingestController, err := ingest.NewController(db, *log)
	if err != nil {
		panic(err)
	}
	metaController, err := meta.NewController(db, *log)
	if err != nil {
		panic(err)
	}

	//TODO: Remove this once Debugging
	fmt.Println("Finished Initialising NewApp")
	return &App{
		IngestController: ingestController,
		MetaController: metaController,
	}
}

func (app *App) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Routes and Controllers
	r.Route("/meta", app.MetaController.Router)
	r.Route("/ingest", app.IngestController.Router)
	return r
}

func (app *App) Shutdown() {
	for _, callback := range app.shutdownCallbacks {
		callback()
	}
}