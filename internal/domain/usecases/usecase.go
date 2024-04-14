package usecases

import (
	"context"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-database-lib/models"
	"log/slog"
)

type UseCase struct {
	requestRepository interfaces.RequestRepository
}

func (u UseCase) Create(ctx context.Context, request models.Request) error {
	slog.InfoContext(ctx, "UseCase.Create",
		slog.String("details", "process started"))

	err := u.requestRepository.Create(ctx, &request)
	if err != nil {
		slog.ErrorContext(ctx, "UseCase.Create",
			slog.String("error", err.Error()))
		return err
	}

	slog.DebugContext(ctx, "UseCase.Create",
		slog.String("details", "process finished"))
	return nil
}

func NewUseCase() interfaces.UseCase {
	return &UseCase{}
}
