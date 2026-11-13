# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener el proyecto completo y arrancable. Guarda los archivos que te devuelva en la
carpeta del repo, respetando exactamente las rutas.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Superficie de practica — NO resuelvas

Estos archivos SON el ejercicio de la persona. No los implementes; deja stubs.

- `internal/domain/interfaces/http/router.go` — La practica caia en codigo que el entry no alcanza: este archivo alcanzable pasa a ser el ejercicio.

## Como saber que terminaste

```bash
go mod tidy && go build ./...
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Perfil
Chapter Backend, Especialidad Desarrollador, Tecnología Go, Senior L2

### Brecha de conocimiento
Construir una API REST en Go con Gin para gestionar tareas, Implementa endpoints REST en Go con Gin para crear, listar y actualizar tareas

### Misión / candidato
Exponer la gestion de tareas por HTTP

### Reto
- Tema: API REST de tareas con Go y Gin
- Seniority: senior-l2
- Tipo: practical
- Título: Gestión de Tareas con API REST en Go y Gin
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Creación de la API — objetivo: Implementar los endpoints básicos para crear y listar tareas. — entregable (NO resolver): Endpoints REST funcionales para crear y listar tareas.
- Fase 2: Actualización de Tareas — objetivo: Implementar el endpoint para actualizar tareas. — entregable (NO resolver): Endpoint REST funcional para actualizar tareas.
- Fase 3: Optimización y Refactorización — objetivo: Optimizar y refactorizar el código para mejorar la legibilidad y mantenibilidad. — entregable (NO resolver): Código refactorizado y optimizado.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (entrega de archivos).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Entrega de archivos
Entregá cada archivo con su ruta completa, respetando exactamente la estructura de rutas
indicada por los marcadores: un bloque por archivo, precedido por su ruta. No generes ZIPs
ni archivos adjuntos.
La entrega debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: Entregá los archivos completos al finalizar. No preguntes si el usuario
quiere recibirlos. Simplemente entregalos; No debes desplegar en el chat el resumen de lo que arreglaste, solo entregá los archivos.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del último archivo entregado
No preguntes si el usuario quiere los archivos: simplemente entregalos siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module github.com/user/project

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
)

// === ARCHIVO: main.go ===
package main

import (
	"context"
	"log"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure"
	httphandler "github.com/user/project/internal/domain/interfaces/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	validate := validator.New()

	taskRepo := infrastructure.NewInMemoryTaskRepository()

	taskService := application.NewTaskService(taskRepo, validate)

	taskHandler := httphandler.NewTaskHandler(taskService)

	httphandler.InitializeRoutes(router, taskHandler)

	server := &stdhttp.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Iniciando servidor en :8080...")
		if err := server.ListenAndServe(); err != nil && err != stdhttp.ErrServerClosed {
			log.Fatalf("Error al iniciar el servidor: %v\n", err)
		}
	}()

	<-quit
	log.Println("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error al apagar el servidor: %v\n", err)
	}
	log.Println("Servidor apagado correctamente")
}

// === ARCHIVO: internal/app/routes.go ===
package app

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	// Sample route - will be replaced by domain code
	router.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sample endpoint",
		})
	})
}

// === ARCHIVO: internal/contracts/types.go ===
package contracts

// Sample contract types - interfaces and models go here

// === ARCHIVO: internal/domain/cmd/api/main.go ===
package main

import (
	"context"
	"log"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure"
	httphandler "github.com/user/project/internal/domain/interfaces/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	validate := validator.New()

	taskRepo := infrastructure.NewInMemoryTaskRepository()

	taskService := application.NewTaskService(taskRepo, validate)

	taskHandler := httphandler.NewTaskHandler(taskService)

	httphandler.InitializeRoutes(router, taskHandler)

	server := &stdhttp.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Iniciando servidor en :8080...")
		if err := server.ListenAndServe(); err != nil && err != stdhttp.ErrServerClosed {
			log.Fatalf("Error al iniciar el servidor: %v\n", err)
		}
	}()

	<-quit
	log.Println("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error al apagar el servidor: %v\n", err)
	}
	log.Println("Servidor apagado correctamente")
}

// === ARCHIVO: internal/domain/task.go ===
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

// === ARCHIVO: internal/domain/application/task_service.go ===
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

// === ARCHIVO: internal/domain/infrastructure/task_repository.go ===
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

// === ARCHIVO: internal/domain/interfaces/http/task_handler.go ===
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/interfaces/http/dto"
)

type TaskHandler struct {
	service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "BAD_REQUEST", ""))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "VALIDATION_ERROR", ""))
		return
	}

	task, err := h.service.CreateTask(req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusConflict, dto.NewErrorResponse(err.Error(), "CONFLICT", ""))
		return
	}

	resp, err := dto.ToTaskResponse(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error(), "INTERNAL_ERROR", ""))
		return
	}

	c.JSON(http.StatusCreated, dto.NewSuccessResponse("Tarea creada exitosamente", resp))
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error(), "INTERNAL_ERROR", ""))
		return
	}

	responses, err := dto.ToTaskResponseList(tasks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error(), "INTERNAL_ERROR", ""))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("Tareas obtenidas exitosamente", responses))
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "BAD_REQUEST", ""))
		return
	}

	task, err := h.service.UpdateTask(id, req.Title, req.Description, req.Status)
	if err != nil {
		if err.Error() == "tarea no encontrada" {
			c.JSON(http.StatusNotFound, dto.NewErrorResponse(err.Error(), "NOT_FOUND", ""))
		} else {
			c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error(), "BAD_REQUEST", ""))
		}
		return
	}

	resp, err := dto.ToTaskResponse(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error(), "INTERNAL_ERROR", ""))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse("Tarea actualizada exitosamente", resp))
}

// === ARCHIVO: internal/domain/interfaces/http/router.go ===
// TODO(reto): archivo stubbeado; implementa los cuerpos.
package http

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes configura las rutas HTTP de la aplicación
func InitializeRoutes(router *gin.Engine, taskHandler *TaskHandler) {
	// Grupo de rutas para tareas
	taskGroup := router.Group("/api/tasks")
	{
		taskGroup.POST("/", taskHandler.CreateTask)
		taskGroup.GET("/", taskHandler.GetAllTasks)
		taskGroup.PUT("/:id", taskHandler.UpdateTask)
	}

	// Middleware global para manejo de errores
	router.Use(errorHandler())
}

// errorHandler maneja errores globales de la aplicación
func errorHandler() gin.HandlerFunc {
  panic("TODO(reto): errorHandler")
}

// === ARCHIVO: internal/domain/interfaces/http/dto/create_task_request.go ===
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

// === ARCHIVO: internal/domain/interfaces/http/dto/update_task_request.go ===
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

// === ARCHIVO: internal/domain/interfaces/http/dto/task_response.go ===
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
```
