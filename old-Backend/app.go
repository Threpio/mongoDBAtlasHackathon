package main

import (
	"context"
	"github.com/threpio/mongoDBAtlasHackathon/backend/db"
	"github.com/threpio/mongoDBAtlasHackathon/backend/types"
)

type App struct {
	Ctx context.Context
	DB  *db.DB

	ServerInfo types.ServerInfo

	shutdownCallbacks []func()
}

func NewApp() (app *App, err error) {
	// TODO: Instantiate properly
	// Add Application etc
	return &App{}, nil
}

func (app *App) Shutdown() {
	for _, callback := range app.shutdownCallbacks {
		callback()
	}
}
