package controllers

import (
	"encoding/json"
	"github.com/PesquisAi/pesquisai-api/internal/config/errortypes"
	"github.com/PesquisAi/pesquisai-api/internal/delivery/dtos"
	"github.com/PesquisAi/pesquisai-api/internal/delivery/validations"
	"github.com/PesquisAi/pesquisai-api/internal/domain/interfaces"
	"github.com/PesquisAi/pesquisai-database-lib/sql/models"
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
		return errortypes.NewValidationException(err.Error())
	}

	err = validations.Validate(&request)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Create",
			slog.String("error", err.Error()))
		return err
	}

	id := uuid.NewString()
	requestModel := models.Request{
		ID:       &id,
		Context:  &request.Context,
		Research: &request.Research,
	}

	err = c.useCase.Create(r.Context(), requestModel)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Create",
			slog.String("error", err.Error()))
		return err
	}

	response := dtos.CreateResponse{RequestId: *requestModel.ID}

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

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) error {
	idStr := r.PathValue("id")
	slog.InfoContext(r.Context(), "Controller.Get",
		slog.String("details", "process started"),
		slog.String("id", idStr))

	id, err := uuid.Parse(idStr)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Get",
			slog.String("error", err.Error()))
		return errortypes.NewValidationException("'id' must be a valid uuid")
	}

	res, err := c.useCase.Get(r.Context(), id)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Get",
			slog.String("error", err.Error()))
		return err
	}

	var getResponse dtos.GetResponse
	getResponse.FromModel(res)

	err = getResponse.WriteHttp(w)
	if err != nil {
		slog.ErrorContext(r.Context(), "Controller.Get",
			slog.String("error", err.Error()))
		return err
	}

	slog.InfoContext(r.Context(), "Controller.Get",
		slog.String("details", "process finished"))
	return nil
}

func NewController(useCase interfaces.UseCase) interfaces.Controller {
	return &Controller{useCase}
}
