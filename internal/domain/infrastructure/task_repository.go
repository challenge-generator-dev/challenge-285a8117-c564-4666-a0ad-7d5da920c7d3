package infrastructure

import (
	"errors"
	"sync"

	"github.com/user/project/internal/domain"
)

// TaskRepository define la interfaz para el repositorio de tareas
type TaskRepository interface {
	Save(task *domain.Task) error
	FindAll() ([]*domain.Task, error)
	FindByID(id string) (*domain.Task, error)
	FindByTitle(title string) (*domain.Task, error)
}

// InMemoryTaskRepository implementa TaskRepository usando memoria
type InMemoryTaskRepository struct {
	tasks map[string]*domain.Task
	mu    sync.RWMutex
}

// NewInMemoryTaskRepository crea una nueva instancia de InMemoryTaskRepository
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

// Save guarda una tarea en el repositorio
func (r *InMemoryTaskRepository) Save(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := task.Validate(); err != nil {
		return err
	}

	r.tasks[task.ID] = task
	return nil
}

// FindAll retorna todas las tareas
func (r *InMemoryTaskRepository) FindAll() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// FindByID busca una tarea por su ID
func (r *InMemoryTaskRepository) FindByID(id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("tarea no encontrada")
	}
	return task, nil
}

// FindByTitle busca una tarea por su título
func (r *InMemoryTaskRepository) FindByTitle(title string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, task := range r.tasks {
		if task.Title == title {
			return task, nil
		}
	}
	return nil, nil
}