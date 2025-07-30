package router

import (
	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/handler"
)

func TaskRouter() *gin.Engine {
	db.InitDB()

	r := gin.Default()
	r.GET("/api/tasks", handler.GetAllTasksHandler)
	return r
}
