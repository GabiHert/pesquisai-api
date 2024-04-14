package dtos

type Request struct {
	Context  string `json:"context" validate:"required,min=100,max=1000"`
	Research string `json:"research" validate:"required,min=10,max=1000"`
}
