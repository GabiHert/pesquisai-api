package queue

import (
	"context"
	"encoding/json"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-database-lib/models"
	"log/slog"
)

type message struct {
	RequestId *string `json:"requestId"`
	Context   *string `json:"context"`
	Research  *string `json:"research"`
}

type AiOrchestratorQueue struct {
	queue interfaces.Queue
}

func (a AiOrchestratorQueue) Publish(ctx context.Context, request *models.Request) error {
	slog.InfoContext(ctx, "useCase.Publish",
		slog.String("details", "Process started"),
		slog.Any("request", *request))

	msg := &message{
		RequestId: request.ID,
		Context:   request.Context,
		Research:  request.Research,
	}
	b, err := json.Marshal(msg)

	err = a.queue.Publish(ctx, b)
	if err != nil {
		slog.ErrorContext(ctx, "useCase.Create",
			slog.String("error", err.Error()))
		return err
	}

	return nil
}

func NewAiOrchestratorQueue(queue interfaces.Queue) interfaces.AiOrchestratorQueue {
	return &AiOrchestratorQueue{queue}
}
