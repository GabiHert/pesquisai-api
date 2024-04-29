package builder

import (
	"context"
	"encoding/json"
	"github.com/GabiHert/pesquisai-api/internal/config/properties"
	"github.com/GabiHert/pesquisai-database-lib/models"
	"log/slog"
)

type message struct {
	RequestId *string `json:"requestId"`
	Context   *string `json:"context"`
	Research  *string `json:"research"`
	Action    *string `json:"action"`
}

func BuildAiOrchestratorMessage(ctx context.Context, request *models.Request) ([]byte, error) {
	slog.InfoContext(ctx, "useCase.Publish",
		slog.String("details", "Process started"),
		slog.Any("request", *request))

	action := properties.AiOrchestratorResearchLocationAction
	msg := &message{
		RequestId: request.ID,
		Context:   request.Context,
		Research:  request.Research,
		Action:    &action,
	}

	return json.Marshal(msg)
}
