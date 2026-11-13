package domain

import (
	"errors"
	"time"
)

// Task representa una tarea en el sistema
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required,min=10,max=500"`
	Status      string    `json:"status" validate:"oneof=pending in-progress completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTask crea una nueva instancia de Task con valores por defecto
func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("el título no puede estar vacío")
	}
	if description == "" {
		return nil, errors.New("la descripción no puede estar vacía")
	}

	return &Task{
		ID:          generateID(),
		Title:       title,
		Description: description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// Update actualiza los campos de la tarea
func (t *Task) Update(title, description, status string) error {
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	if status != "" {
		t.Status = status
	}
	t.UpdatedAt = time.Now()
	return nil
}

// Validate valida los campos de la tarea según las reglas de negocio
func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("el título es obligatorio")
	}
	if len(t.Title) < 3 || len(t.Title) > 100 {
		return errors.New("el título debe tener entre 3 y 100 caracteres")
	}
	if t.Description == "" {
		return errors.New("la descripción es obligatoria")
	}
	if len(t.Description) < 10 || len(t.Description) > 500 {
		return errors.New("la descripción debe tener entre 10 y 500 caracteres")
	}
	if t.Status != "pending" && t.Status != "in-progress" && t.Status != "completed" {
		return errors.New("el estado debe ser pending, in-progress o completed")
	}
	return nil
}

// generateID genera un ID único para la tarea (simulación)
func generateID() string {
	return "task-" + time.Now().Format("20060102-150405.999999999")
}