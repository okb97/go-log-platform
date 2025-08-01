package repository

import (
	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/model"
)

func GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task

	if err := db.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(task *model.Task) error {
	return db.DB.Create(&task).Error
}
