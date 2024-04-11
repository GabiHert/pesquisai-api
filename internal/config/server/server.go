package server

import (
	"github.com/GabiHert/pesquisai-api/internal/config/routes"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"net/http"
)

func Serve(mux *http.ServeMux, controller interfaces.Controller) error {
	routes.Init(mux, controller)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return server.ListenAndServe()
}
