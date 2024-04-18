package interfaces

import (
	"context"
	"github.com/GabiHert/pesquisai-database-lib/models"
)

type AiOrchestratorQueue interface {
	Publish(ctx context.Context, request *models.Request) (err error)
}
