package controllers

import (
	"encoding/json"
	"github.com/GabiHert/pesquisai-api/internal/delivery/dtos"
	"github.com/GabiHert/pesquisai-api/internal/delivery/validations"
	"github.com/GabiHert/pesquisai-api/internal/domain/interfaces"
	"github.com/GabiHert/pesquisai-database-lib/models"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type Controller struct {
	useCase interfaces.UseCase
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) error {
	slog.InfoContext(r.Context(), "Controller.Create",
		slog.String("details", "process started"))

	var request dtos.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Create",
			slog.String("error", err.Error()))
		return err
	}

	err = validations.Validate(&request)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Create",
			slog.String("error", err.Error()))
		return err
	}

	requestModel := models.Request{
		ID:       uuid.NewString(),
		Context:  request.Context,
		Research: request.Research,
	}

	err = c.useCase.Create(r.Context(), requestModel)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Create",
			slog.String("error", err.Error()))
		return err
	}

	response := dtos.Response{RequestId: requestModel.ID}

	err = response.WriteHttp(w)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Create",
			slog.String("error", err.Error()))
		return err
	}

	slog.InfoContext(r.Context(), "Controller.Create",
		slog.String("details", "process finished"))
	return nil
}

func NewController(useCase interfaces.UseCase) interfaces.Controller {
	return &Controller{useCase}
}
