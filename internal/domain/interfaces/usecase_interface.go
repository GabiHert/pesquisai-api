package interfaces

import (
	"context"
	"github.com/GabiHert/pesquisai-database-lib/models"
)

type UseCase interface {
	Create(ctx context.Context, request models.Request) error
}
