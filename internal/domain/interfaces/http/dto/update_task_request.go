package dto

// TODO(reto): Implementar la validación del DTO para actualización de tareas.
// El DTO debe incluir los campos: Title, Description y Status.
// Validaciones requeridas:
// - Title: requerido, mínimo 3 caracteres, máximo 100
// - Description: mínimo 5 caracteres, máximo 500 (opcional)
// - Status: requerido, debe ser uno de: "pending", "in_progress", "completed"
type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"min=5,max=500"`
	Status      string `json:"status" binding:"required"`
}