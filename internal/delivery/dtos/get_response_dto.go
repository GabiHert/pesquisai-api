package dtos

import (
	"encoding/json"
	"github.com/PesquisAi/pesquisai-database-lib/models"
	"net/http"
	"time"
)

type Research struct {
	Id        *string    `json:"id"`
	Title     *string    `json:"title,omitempty"`
	Link      *string    `json:"link,omitempty"`
	Status    *string    `json:"status"`
	Summary   *string    `json:"summary,omitempty"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type GetResponse struct {
	Id                      *string     `json:"id"`
	Context                 *string     `json:"context"`
	Research                *string     `json:"research"`
	TotalResearches         *int        `json:"total_researches"`
	TotalFinishedResearches *int        `json:"total_finished_researches"`
	Status                  *string     `json:"status"`
	Overall                 *string     `json:"overall,omitempty"`
	CreatedAt               *time.Time  `json:"created_at"`
	UpdatedAt               *time.Time  `json:"updated_at,omitempty"`
	Researches              *[]Research `json:"researches"`
}

func (r *GetResponse) FromModel(model *models.Request) {
	r.Id = model.ID
	r.Context = model.Context
	r.TotalResearches = model.TotalResearches
	r.TotalFinishedResearches = model.TotalFinishedResearches
	r.Research = model.Research
	r.Overall = model.Overall
	r.Status = model.Status
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	r.Researches = &[]Research{}
	*r.Researches = make([]Research, len(model.Researches))
	for i, research := range model.Researches {
		(*r.Researches)[i] = Research{
			Id:        research.ID,
			Title:     research.Title,
			Link:      research.Link,
			Status:    research.Status,
			Summary:   research.Summary,
			CreatedAt: research.CreatedAt,
			UpdatedAt: research.UpdatedAt,
		}
	}
}

func (r *GetResponse) WriteHttp(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(r)
}
