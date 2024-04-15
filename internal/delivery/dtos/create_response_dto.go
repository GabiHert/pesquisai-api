package dtos

import (
	"encoding/json"
	"net/http"
)

type CreateResponse struct {
	RequestId string `json:"id"`
}

func (r *CreateResponse) WriteHttp(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(r)
}
