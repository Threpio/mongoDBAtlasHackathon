package main

import (
	"fmt"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/app"
	"github.com/threpio/mongoDBAtlasHackathon/backend2/db"
	"os"
	"os/signal"
)

// cli.go Is where the actual main.go function is found.

func StartServer() {
	DB, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	srv := app.NewApp(*DB)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	srv.Serve()

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
