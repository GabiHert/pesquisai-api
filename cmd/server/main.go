package main

import (
	"github.com/PesquisAi/pesquisai-api/internal/clients/connections"
	"github.com/PesquisAi/pesquisai-api/internal/config/injector"
	"github.com/PesquisAi/pesquisai-api/internal/config/server"
)

func main() {
	deps := injector.NewDependencies()
	var err error

	if err = connections.Connect(deps); err != nil {
		panic(err)
	}

	if err = server.Serve(deps.Mux, deps.Controller); err != nil {
		panic(err)
	}
}
