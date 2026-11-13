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