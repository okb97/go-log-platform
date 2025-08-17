package repository

import (
	"github.com/okb97/go-log-platform/internal/model"
)

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task *model.Task) error
	DeleteTask(id uint) error
	UpdateTask(task *model.Task) error
}
