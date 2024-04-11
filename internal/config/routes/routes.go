package routes

import (
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"net/http"
)

func Init(mux *http.ServeMux, controller interfaces.Controller) {
	mux.HandleFunc("POST /pesquisai", controller.Create)
}
