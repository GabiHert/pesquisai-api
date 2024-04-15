package server

import (
	"fmt"
	"github.com/GabiHert/pesquisai-api/internal/config/routes"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"log/slog"
	"net/http"
)

func Serve(mux *http.ServeMux, controller interfaces.Controller) error {
	routes.Init(mux, controller)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	slog.Info(fmt.Sprintf("Starting server at port '%s'", server.Addr))
	return server.ListenAndServe()
}
