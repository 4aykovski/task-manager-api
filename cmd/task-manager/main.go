package main

import (
	"github.com/4aykovski/task-manager-api/internal/config"
	"github.com/4aykovski/task-manager-api/pkg/libs/logger/slogHelper"
)

func main() {

	cfg := config.MustLoad()

	log := slogHelper.SetupLogger(cfg.Env)

	log.Debug("123")

}
