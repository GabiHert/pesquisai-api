package usecases

import (
	"context"
	"fmt"
	"github.com/GabiHert/pesquisai-api/internal/config/errortypes"
	"github.com/GabiHert/pesquisai-api/internal/domain/builder"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-database-lib/models"
	"github.com/google/uuid"
	"log/slog"
	"strings"
)

type UseCase struct {
	requestRepository   interfaces.RequestRepository
	aiOrchestratorQueue interfaces.Queue
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

	b, err := builder.BuildAiOrchestratorMessage(ctx, &request)
	if err != nil {
		slog.ErrorContext(ctx, "useCase.Create",
			slog.String("error", err.Error()))
		return err
	}

	err = u.aiOrchestratorQueue.Publish(ctx, b)
	if err != nil {
		slog.ErrorContext(ctx, "useCase.Create",
			slog.String("error", err.Error()))
		return err
	}

	slog.DebugContext(ctx, "useCase.Create",
		slog.String("details", "process finished"))
	return nil
}

func (u UseCase) Get(ctx context.Context, id uuid.UUID) (*models.Request, error) {
	slog.InfoContext(ctx, "useCase.Get",
		slog.String("details", "process started"))

	res, err := u.requestRepository.GetWithResearches(ctx, id.String())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = errortypes.NewNotFoundException(fmt.Sprintf("Could not find register with id '%s'", id.String()))
		}
		slog.ErrorContext(ctx, "useCase.Get",
			slog.String("error", err.Error()))
		return nil, err
	}

	slog.DebugContext(ctx, "useCase.Get",
		slog.String("details", "process finished"))
	return res, nil
}

func NewUseCase(requestRepository interfaces.RequestRepository, aiOrchestratorQueue interfaces.Queue) interfaces.UseCase {
	return &UseCase{
		requestRepository:   requestRepository,
		aiOrchestratorQueue: aiOrchestratorQueue,
	}
}
