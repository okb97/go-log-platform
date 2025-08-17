package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/internal/model"
	"github.com/okb97/go-log-platform/internal/service"
)

type TaskHandler struct {
	service service.TaskServiceInterface
}

func NewTaskHandler(service service.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetAllTasksHandler(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "タスク一覧の取得に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTaskHandler(c *gin.Context) {
	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) DeleteTaskHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}
	if err := h.service.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "タスクの削除に失敗しました"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = uint(id)

	if err := h.service.UpdateTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update task"})
		return
	}

	c.Status(http.StatusNoContent)

}
