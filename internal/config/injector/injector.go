package injector

import (
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-database-lib/connection"
	"github.com/GabiHert/pesquisai-database-lib/repositories"
	"gorm.io/gorm"
	"net/http"
)

type Dependencies struct {
	Mux               *http.ServeMux
	Controller        interfaces.Controller
	RequestRepository interfaces.RequestRepository
	Connection        *connection.Connection
}

func (d *Dependencies) Inject() *Dependencies {
	if d.Connection == nil {
		d.Connection = &connection.Connection{DB: &gorm.DB{}}
	}

	if d.Mux == nil {
		d.Mux = http.NewServeMux()
	}

	if d.RequestRepository == nil {
		d.RequestRepository = &repositories.Request{Connection: d.Connection}
	}
	return d
}
