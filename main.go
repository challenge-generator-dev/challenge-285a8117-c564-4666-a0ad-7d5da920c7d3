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