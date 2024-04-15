package routes

import (
	"errors"
	"github.com/GabiHert/pesquisai-api/internal/config/errortypes"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-errors-lib/exceptions"
	"log/slog"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	var customError *exceptions.Error
	ok := errors.As(err, &customError)
	if !ok {
		customError = errortypes.NewUnknownException(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	err2 := customError.WriteHttp(w)
	if err2 != nil {
		slog.Error("routes.handleError", "errorMessage", err2.Error())
		panic(err2)
	}
}

func Init(mux *http.ServeMux, controller interfaces.Controller) {
	mux.HandleFunc("POST /v1/pesquisai", func(writer http.ResponseWriter, request *http.Request) {
		handleError(writer, controller.Create(writer, request))
	})

	mux.HandleFunc("GET /v1/pesquisai/{id}", func(writer http.ResponseWriter, request *http.Request) {
		handleError(writer, controller.Get(writer, request))
	})
}
