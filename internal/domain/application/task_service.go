package application

import (
	"github.com/go-playground/validator/v10"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/infrastructure"
)

// TaskService define la interfaz para el servicio de tareas
type TaskService struct {
	repo      infrastructure.TaskRepository
	validator *validator.Validate
}

// NewTaskService crea una nueva instancia de TaskService
func NewTaskService(repo infrastructure.TaskRepository, validator *validator.Validate) *TaskService {
	return &TaskService{
		repo:      repo,
		validator: validator,
	}
}

// CreateTask crea una nueva tarea
func (s *TaskService) CreateTask(title, description string) (*domain.Task, error) {
	panic("TODO(reto): CreateTask")
}

// GetAllTasks retorna todas las tareas
func (s *TaskService) GetAllTasks() ([]*domain.Task, error) {
	panic("TODO(reto): GetAllTasks")
}

// UpdateTask actualiza una tarea existente
func (s *TaskService) UpdateTask(id, title, description, status string) (*domain.Task, error) {
	panic("TODO(reto): UpdateTask")
}