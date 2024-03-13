package main

import (
	"fmt"
	"net/http"

	"github.com/4aykovski/task-manager-api/internal/config"
	v1 "github.com/4aykovski/task-manager-api/internal/net/v1"
	"github.com/4aykovski/task-manager-api/pkg/libs/logger/slogHelper"
)

func main() {

	cfg := config.MustLoad()

	log := slogHelper.SetupLogger(cfg.Env)

	mux := v1.NewMux(log)

	server := http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      mux,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	log.Info(fmt.Sprintf("starting server on %s", cfg.HTTPServer.Address))

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	log.Error("server stopped")
}
