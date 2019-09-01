package main

import (
	"net/http"
	"resumator-backend/config"
	"resumator-backend/logger"
	"resumator-backend/routes"
	"sync"
	"time"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger
var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(1)
	log = logger.GetLogger()
	log.Info("starting the Resumator Backend...")
	config.Initialize()
	log.Infof("config linkedIn url %s", config.GetLinkedInAccessTokenURL())
	go startServer()
	waitGroup.Wait()
}

func startServer() {
	router := routes.InitializeRouter()

	Server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Info("started listening at 8000")
	Server.ListenAndServe()
}
