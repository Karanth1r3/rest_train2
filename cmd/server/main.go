package main

import (
	"log/slog"
	"os"

	apiserver "github.com/Karanth1r3/rest-train-2/internal/app/api_server"
	"github.com/Karanth1r3/rest-train-2/internal/config"
	"github.com/Karanth1r3/rest-train-2/internal/utils/slg"
)

var (
	configPath = "configs/config.json"
)

func main() {

	// Initializing logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))

	//Initializing config
	cfg, err := config.ReadConfig(configPath)
	if err != nil {
		slg.LogErrorFatal(logger, "could not initialize config: ", err)
	}
	_ = cfg
	serv := apiserver.New()

	if err := serv.Start(); err != nil {
		//log.Fatal()
		slg.LogErrorFatal(logger, "could not start apiserver:", err)
	}
}
