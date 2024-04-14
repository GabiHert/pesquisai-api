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
	slog.InfoContext(ctx, "useCase.Create",
		slog.String("details", "process started"))

	err := u.requestRepository.Create(ctx, &request)
	if err != nil {
		slog.ErrorContext(ctx, "useCase.Create",
			slog.String("error", err.Error()))
		return err
	}

	slog.DebugContext(ctx, "useCase.Create",
		slog.String("details", "process finished"))
	return nil
}

func NewUseCase(requestRepository interfaces.RequestRepository) interfaces.UseCase {
	return &UseCase{
		requestRepository: requestRepository,
	}
}
