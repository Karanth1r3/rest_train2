package apiserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	apiserver "github.com/Karanth1r3/rest-train-2/internal/app/api_server"
	"github.com/Karanth1r3/rest-train-2/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {

	cfg := config.Config{
		HTTPServer: config.HTTPServer{":8081"},
		Logger:     config.Logger{"info"},
	}

	serv := apiserver.New(&cfg)

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	serv.HandleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
