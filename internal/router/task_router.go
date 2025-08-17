package router

import (
	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/internal/handler"
	"github.com/okb97/go-log-platform/internal/repository"
	"github.com/okb97/go-log-platform/internal/service"
)

func TaskRouter() *gin.Engine {

	r := gin.Default()
	taskRepo := repository.NewGormTaskRepository()
	taskService := service.NewTaskService(taskRepo)

	handler := handler.NewTaskHandler(taskService)

	r.GET("/api/tasks", handler.GetAllTasksHandler)
	r.POST("/api/task", handler.CreateTaskHandler)
	r.DELETE("/api/tasks/:id", handler.DeleteTaskHandler)
	r.PUT("/api/tasks/:id", handler.UpdateTaskHandler)
	return r
}
