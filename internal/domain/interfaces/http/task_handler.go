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