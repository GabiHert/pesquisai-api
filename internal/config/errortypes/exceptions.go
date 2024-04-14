package errortypes

import (
	"github.com/GabiHert/perguntai-errors-lib/exceptions"
	"net/http"
)

func NewUnknownException(message string) *exceptions.Error {
	return &exceptions.Error{Message: message, ErrorType: exceptions.ErrorType{
		Code:           "PAPI01",
		Type:           "Unknown",
		HttpStatusCode: http.StatusInternalServerError,
	}}
}
