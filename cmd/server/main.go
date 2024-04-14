package main

import (
	"github.com/GabiHert/pesquisai-api/internal/config/injector"
	"github.com/GabiHert/pesquisai-api/internal/config/server"
	"github.com/GabiHert/pesquisai-database-lib/connection"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	deps := injector.Dependencies{}
	deps.Inject()
	err := deps.Connection.Connect(connection.Config{
		User: "postgres",
		Host: "localhost",
		Psw:  "postgres",
		Name: "postgres",
		Port: "5432",
		GormConfig: gorm.Config{
			NamingStrategy: schema.NamingStrategy{TablePrefix: "pesquisai."},
		},
	})
	if err != nil {
		panic(err)
	}

	if err := server.Serve(deps.Mux, deps.Controller); err != nil {
		panic(err)
	}
}
