package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/user/project/internal/domain"
)

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=5,max=500"`
}

func (r *CreateTaskRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *CreateTaskRequest) ToDomain() (*domain.Task, error) {
	return domain.NewTask(r.Title, r.Description)
}