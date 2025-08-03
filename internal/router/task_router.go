package router

import (
	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/internal/handler"
)

func TaskRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/api/tasks", handler.GetAllTasksHandler)
	r.POST("/api/task", handler.CreateTaskHandler)
	r.DELETE("/api/tasks/:id", handler.DeleteTaskHandler)
	return r
}
