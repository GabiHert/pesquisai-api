package interfaces

import (
	"context"
	"github.com/PesquisAi/pesquisai-database-lib/models"
	"github.com/google/uuid"
)

type UseCase interface {
	Create(ctx context.Context, request models.Request) error
	Get(ctx context.Context, id uuid.UUID) (*models.Request, error)
}
