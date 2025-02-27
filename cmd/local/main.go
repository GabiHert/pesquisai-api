package main

import (
	"github.com/PesquisAi/pesquisai-api/internal/clients/connections"
	"github.com/PesquisAi/pesquisai-api/internal/config/injector"
	"github.com/PesquisAi/pesquisai-api/internal/config/server"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		panic(err)
	}
	deps := injector.NewDependencies()

	if err = connections.Connect(deps); err != nil {
		panic(err)
	}

	if err = server.Serve(deps.Mux, deps.Controller); err != nil {
		panic(err)
	}
}
