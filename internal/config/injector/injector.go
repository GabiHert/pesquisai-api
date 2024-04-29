package injector

import (
	"github.com/GabiHert/pesquisai-api/internal/config/properties"
	"github.com/GabiHert/pesquisai-api/internal/delivery/controllers"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-api/internal/domain/usecases"
	"github.com/GabiHert/pesquisai-database-lib/connection"
	"github.com/GabiHert/pesquisai-database-lib/repositories"
	"github.com/GabiHert/pesquisai-rabbitmq-lib/rabbitmq"
	"gorm.io/gorm"
	"net/http"
)

type Dependencies struct {
	Mux                 *http.ServeMux
	Controller          interfaces.Controller
	RequestRepository   interfaces.RequestRepository
	DatabaseConnection  *connection.Connection
	QueueConnection     *rabbitmq.Connection
	UseCase             interfaces.UseCase
	AiOrchestratorQueue interfaces.Queue
}

func (d *Dependencies) Inject() *Dependencies {
	if d.DatabaseConnection == nil {
		d.DatabaseConnection = &connection.Connection{DB: &gorm.DB{}}
	}

	if d.Mux == nil {
		d.Mux = http.NewServeMux()
	}

	if d.RequestRepository == nil {
		d.RequestRepository = &repositories.Request{Connection: d.DatabaseConnection}
	}

	if d.QueueConnection == nil {
		d.QueueConnection = &rabbitmq.Connection{}
	}

	if d.AiOrchestratorQueue == nil {
		d.AiOrchestratorQueue = rabbitmq.NewQueue(d.QueueConnection,
			properties.AiOrchestratorQueueName,
			rabbitmq.CONTENT_TYPE_JSON,
			properties.CreateQueueIfNX())
	}

	if d.UseCase == nil {
		d.UseCase = usecases.NewUseCase(d.RequestRepository, d.AiOrchestratorQueue)
	}

	if d.Controller == nil {
		d.Controller = controllers.NewController(d.UseCase)
	}
	return d
}

func NewDependencies() *Dependencies {
	deps := &Dependencies{}
	deps.Inject()
	return deps
}
