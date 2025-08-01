package service

import (
	"github.com/okb97/go-log-platform/internal/model"
	"github.com/okb97/go-log-platform/internal/repository"
)

func GetAllTasks() ([]model.Task, error) {
	return repository.GetAllTasks()
}

func CreateTask(task *model.Task) error {
	return repository.CreateTask(task)
}
