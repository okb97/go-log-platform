package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/internal/model"
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

func CreateTaskHandler(c *gin.Context) {
	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}
