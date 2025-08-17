package service

import (
	"github.com/okb97/go-log-platform/internal/model"
	"github.com/okb97/go-log-platform/internal/repository"
)

type TaskServiceInterface interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task *model.Task) error
	DeleteTask(id uint) error
	UpdateTask(task *model.Task) error
}

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) CreateTask(task *model.Task) error {
	return s.repo.CreateTask(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) UpdateTask(task *model.Task) error {
	return s.repo.UpdateTask(task)
}
