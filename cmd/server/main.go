package main

import (
	"log"

	apiserver "github.com/Karanth1r3/rest-train-2/internal/app/api_server"
	"github.com/Karanth1r3/rest-train-2/internal/config"
	"github.com/Karanth1r3/rest-train-2/internal/utils/slg"
)

const (
	infoLevel  = "info"
	debugLevel = "debug"
)

var (
	configPath = "configs/config.json"
)

func main() {

	// Initializing logger

	//Initializing config
	cfg, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	serv := apiserver.New(cfg)

	if err := serv.Start(); err != nil {
		//log.Fatal()
		slg.LogErrorFatal(serv.Logger, "could not start apiserver:", err)
	}
}

// For experimental reasons this one is put into app.apiserver.apiserver.go (looks like it's a bad idea - it's fields are public to be able to be called from main)
/*
func setupLogger(cfg config.Config) *slog.Logger {
	var log *slog.Logger
	switch cfg.Logger.LogLevel {
	case debugLevel:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case infoLevel:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
*/
