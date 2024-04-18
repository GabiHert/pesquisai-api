package injector

import (
	"github.com/GabiHert/pesquisai-api/internal/delivery/controllers"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-api/internal/domain/queue"
	"github.com/GabiHert/pesquisai-api/internal/domain/usecases"
	"github.com/GabiHert/pesquisai-database-lib/connection"
	"github.com/GabiHert/pesquisai-database-lib/repositories"
	rabbitmqConnection "github.com/GabiHert/pesquisai-rabbitmq-lib/connection"
	rabbitmq "github.com/GabiHert/pesquisai-rabbitmq-lib/queue"
	"gorm.io/gorm"
	"net/http"
)

type Dependencies struct {
	Mux                 *http.ServeMux
	Controller          interfaces.Controller
	RequestRepository   interfaces.RequestRepository
	Connection          *connection.Connection
	QueueConnection     *rabbitmqConnection.Connection
	UseCase             interfaces.UseCase
	Queue               interfaces.Queue
	AiOrchestratorQueue interfaces.AiOrchestratorQueue
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

	if d.QueueConnection == nil {
		d.QueueConnection = &rabbitmqConnection.Connection{}
	}

	if d.Queue == nil {
		d.Queue = rabbitmq.NewQueue(d.QueueConnection, "", "", false)
	}

	if d.AiOrchestratorQueue == nil {
		d.AiOrchestratorQueue = queue.NewAiOrchestratorQueue(d.Queue)
	}

	if d.UseCase == nil {
		d.UseCase = usecases.NewUseCase(d.RequestRepository, d.AiOrchestratorQueue)
	}

	if d.Controller == nil {
		d.Controller = controllers.NewController(d.UseCase)
	}
	return d
}
