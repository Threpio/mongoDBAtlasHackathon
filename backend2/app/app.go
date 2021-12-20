package app

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/ingest"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/logger"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/meta"
	"net/http"
)

type App struct {
	Ctx context.Context
	shutdownCallbacks []func()

	IngestController *ingest.Controller
	MetaController   *meta.Controller
}

func NewApp(db db.DB) *App {

	log := logger.New("DEBUG")

	ingestController, err := ingest.NewController(db, *log)
	if err != nil {
		panic(err)
	}
	metaController, err := meta.NewController(db, *log)
	if err != nil {
		panic(err)
	}

	return &App{
		IngestController: ingestController,
		MetaController: metaController,
	}
}

func (app *App) Serve() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Routes and Controllers
	r.Route("/meta", app.MetaController.Router)
	r.Route("/ingest", app.IngestController.Router)
	http.ListenAndServe(":8080", r)
}

func (app *App) Shutdown() {
	for _, callback := range app.shutdownCallbacks {
		callback()
	}
}