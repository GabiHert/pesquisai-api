package routes

import (
	"errors"
	"github.com/GabiHert/perguntai-errors-lib/exceptions"
	"github.com/GabiHert/pesquisai-api/internal/config/errortypes"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"log/slog"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	var customError *exceptions.Error
	ok := errors.As(err, customError)
	if !ok {
		customError = errortypes.NewUnknownException(err.Error())
	}

	err2 := customError.WriteHttp(w)
	if err2 != nil {
		slog.Error("routes.handleError", "errorMessage", err2.Error())
		panic(err2)
	}
}

func Init(mux *http.ServeMux, controller interfaces.Controller) {
	mux.HandleFunc("POST /pesquisai", func(writer http.ResponseWriter, request *http.Request) {
		handleError(writer, controller.Create(writer, request))
	})
}
