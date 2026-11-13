package dto

import (
	"time"

	"github.com/user/project/internal/domain"
)

// TaskResponse representa la estructura de respuesta para una tarea.
// Se utiliza para serializar los datos de una tarea en las respuestas HTTP.
type TaskResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToTaskResponse convierte una entidad Task del dominio en un DTO de respuesta.
// Esta conversión aísla la estructura interna del dominio de la representación externa.
func ToTaskResponse(task *domain.Task) (*TaskResponse, error) {
	if task == nil {
		return nil, nil
	}

	if err := task.Validate(); err != nil {
		return nil, err
	}

	return &TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

// ToTaskResponseList convierte una lista de entidades Task en una lista de DTOs de respuesta.
func ToTaskResponseList(tasks []*domain.Task) ([]*TaskResponse, error) {
	responses := make([]*TaskResponse, 0, len(tasks))

	for _, task := range tasks {
		resp, err := ToTaskResponse(task)
		if err != nil {
			return nil, err
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

// TaskListResponse representa la estructura de respuesta para una lista de tareas.
type TaskListResponse struct {
	Tasks    []*TaskResponse `json:"tasks"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"page_size"`
}

// NewTaskListResponse crea una nueva respuesta de lista de tareas con metadatos de paginación.
func NewTaskListResponse(tasks []*TaskResponse, total, page, pageSize int) *TaskListResponse {
	return &TaskListResponse{
		Tasks:    tasks,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

// ErrorResponse representa la estructura de respuesta para errores.
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}

// NewErrorResponse crea una nueva respuesta de error.
func NewErrorResponse(message string, code string, details string) *ErrorResponse {
	return &ErrorResponse{
		Error:   message,
		Code:    code,
		Details: details,
	}
}

// SuccessResponse representa la estructura de respuesta genérica para operaciones exitosas.
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewSuccessResponse crea una nueva respuesta de éxito.
func NewSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Data:    data,
	}
}