package interfaces

import "net/http"

type Controller interface {
	Create(w http.ResponseWriter, r *http.Request) error
	Get(w http.ResponseWriter, r *http.Request) error
}
