package apiserver

// APIServer ...
type APIServer struct{}

// CTOR
func New() *APIServer {
	return &APIServer{}
}

// Laynch api server
func (s *APIServer) Start() error {
	return nil
}
