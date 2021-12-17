package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func server() {
	app, err := NewApp()
	if err != nil {
		panic(err)
	}

	// Subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	// Serve stuff here

	// TODO: This?
	apiHandler := NewHandler(&app)
	apiAddress := fmt.Sprintf("%s:%s", viper.GetString("core-ip"), viper.GetString("core-port"))

	var srv *http.Server
	srv = &http.Server{
		Addr:         apiAddress,
		Handler:      apiHandler,
		WriteTimeout: 1 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}

	// Wait for SIGINT
	<-stopChan
	fmt.Print("\r") // overwrite the ^C
	log.Println("Shutting down server...(Ctrl+c to force)")
	defer func() {
		// Implement these things
		//app.Logger.Info("App is stopping...")
		//app.Shutdown()
		//app.Logger.Info("App stopped")
	}()
}

func runServer() error {
	server()
	return nil
}