package router

import (
	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/handler"
	"github.com/okb97/go-log-platform/internal/repository"
)

func TaskRouter() *gin.Engine {
	db.InitDB()
	repository.SeedTestData()

	r := gin.Default()
	r.GET("/api/tasks", handler.GetAllTasksHandler)
	return r
}
