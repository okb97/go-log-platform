package router

import (
	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/internal/handler"
)

func TaskRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/api/tasks", handler.GetAllTasksHandler)
	return r
}
