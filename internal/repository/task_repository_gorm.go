package repository

import (
	"fmt"

	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/model"
)

type GormTaskRepository struct{}

func NewGormTaskRepository() *GormTaskRepository {
	return &GormTaskRepository{}
}

func (r *GormTaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task

	if err := db.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *GormTaskRepository) CreateTask(task *model.Task) error {
	return db.DB.Create(&task).Error
}

func (r *GormTaskRepository) DeleteTask(id uint) error {
	result := db.DB.Delete(&model.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("task with ID %d not found", id)
	}
	return nil
}

func (r *GormTaskRepository) UpdateTask(task *model.Task) error {
	return db.DB.Model(&model.Task{}).
		Where("id = ?", task.ID).
		Omit("CreatedAt").
		Updates(task).Error
}
