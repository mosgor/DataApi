package main

import (
	"SourceManager/internal/config"
	"SourceManager/internal/http_server"
	logs "SourceManager/internal/logger"
	_ "net"
)

func main() {
	cfg := config.MustLoad()

	log := logs.SetupLogger(cfg.Env)

	log.Info("Logger is up")

	server := http_server.CreateServer()

	err := server.ListenAndServe()
	if err != nil {
		log.Error(err.Error())
	}

	// TODO: дополнить запросы данными с дб
}
