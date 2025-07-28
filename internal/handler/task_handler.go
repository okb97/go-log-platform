package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/internal/service"
)

func GetAllTasksHandler(c *gin.Context) {
	tasks, err := service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "タスク一覧の取得に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
