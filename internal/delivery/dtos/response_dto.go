package dtos

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	RequestId string `json:"request-id"`
}

func (r *Response) WriteHttp(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(r)
}
