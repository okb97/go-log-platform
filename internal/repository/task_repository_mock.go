package repository

import (
	"fmt"

	"github.com/okb97/go-log-platform/internal/model"
)

type MockTaskRepository struct {
	Tasks []model.Task
}

func NewMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{
		Tasks: []model.Task{},
	}
}

func (m *MockTaskRepository) GetAllTasks() ([]model.Task, error) {
	return m.Tasks, nil
}

func (m *MockTaskRepository) CreateTask(task *model.Task) error {
	task.ID = uint(len(m.Tasks) + 1)
	m.Tasks = append(m.Tasks, *task)
	return nil
}

func (m *MockTaskRepository) DeleteTask(id uint) error {
	for i, t := range m.Tasks {
		if t.ID == id {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (m *MockTaskRepository) UpdateTask(task *model.Task) error {
	for i, t := range m.Tasks {
		if t.ID == task.ID {
			m.Tasks[i] = *task
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", task.ID)
}
