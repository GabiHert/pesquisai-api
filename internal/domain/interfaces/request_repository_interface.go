package interfaces

import (
	"github.com/GabiHert/pesquisai-database-lib/models"
)

type RequestRepository interface {
	Create(request *models.Request) error
}
