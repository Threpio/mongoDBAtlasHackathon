package main

import (
	"errors"
	"fmt"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/app"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/db"
	"net/http"
	"os"
	"os/signal"
)

// cli.go Is where the actual main.go function is found.

const (
	PORT = "8080"
)

// StartServer is the main function of this application.
// It returns an error that the application throws if it panics or shutdown gracefully.
// It also returns an error if the application is unable to start.
func StartServer() error {
	DB, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	app := app.NewApp(*DB)

	//TODO: Remove this once Debugging
	fmt.Println("Initialising stopChannel")
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	//TODO: Remove this once Debugging
	fmt.Println("Finished Initialising stopChannel")

	fmt.Printf("Starting server on port %s\n", PORT)

	var srv *http.Server
	go func() {
		srv = &http.Server{
			Addr: ":" + PORT,
			Handler: app.Router(),
		}
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<- stopChan
	// Do app.Stop
	fmt.Println("\r")
	fmt.Println("Stopping server...")
	defer func() {
		// Implement these things
		//app.Logger.Info("App is stopping...")
		//app.Shutdown()
		//app.Logger.Info("App stopped")
	}()
}
