package apiserver

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Karanth1r3/rest-train-2/internal/config"
	"github.com/gorilla/mux"
)

const (
	infoLevel  = "info"
	debugLevel = "debug"
)

// APIServer ...
type APIServer struct {
	Config *config.Config
	Logger *slog.Logger
	router *mux.Router
}

// CTOR
func New(cfg *config.Config) *APIServer {
	return &APIServer{
		Config: cfg,
		Logger: setupLogger(*cfg),
		router: mux.NewRouter(),
	}
}

// Laynch api server
func (s *APIServer) Start() error {
	// Initializing logger
	s.Logger.Info("info messages are enabled")
	s.Logger.Debug("debug messages are enabled")
	s.Logger.Info("starting api server")

	s.Logger.Info("configuring router")
	s.setupRouter()

	s.Logger.Info("trying to launch server")
	return http.ListenAndServe(s.Config.Address, s.router)
}

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

func (s *APIServer) setupRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello"))
	}
}
