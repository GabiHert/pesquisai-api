package main

import (
	"github.com/GabiHert/pesquisai-api/internal/config/injector"
	"github.com/GabiHert/pesquisai-api/internal/config/server"
)

func main() {
	deps := injector.Dependencies{}
	deps.Inject()

	if err := server.Serve(deps.Mux, deps.Controller); err != nil {
		panic(err)
	}
}
