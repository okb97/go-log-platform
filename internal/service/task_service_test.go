package service

import (
	"testing"

	"github.com/okb97/go-log-platform/internal/model"
	"github.com/okb97/go-log-platform/internal/repository"
)

func TestGetAllTasks(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	mockRepo.Tasks = []model.Task{
		{ID: 1, Title: "買い物", Status: "pending"},
		{ID: 2, Title: "Go学習", Status: "completed"},
	}
	service := NewTaskService(mockRepo)

	tasks, err := service.GetAllTasks()
	if err != nil {
		t.Fatalf("GetAllTasks() returned error: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

}

func TestCreateTask(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	service := NewTaskService(mockRepo)

	task := &model.Task{ID: 1, Title: "テストタスク", Status: "pending"}
	if err := service.CreateTask(task); err != nil {
		t.Fatalf("CreateTask() returned error: %v", err)
	}

	if len(mockRepo.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(mockRepo.Tasks))
	}
	if mockRepo.Tasks[0].Title != "テストタスク" {
		t.Errorf("Expected Title 'テストタスク', got %q", mockRepo.Tasks[0].Title)
	}
}

func TestDeleteTask(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	task := &model.Task{ID: 1, Title: "旧タスク", Status: "pending"}
	mockRepo.Tasks = append(mockRepo.Tasks, *task)
	service := NewTaskService(mockRepo)

	task.Title = "更新後タスク"
	if err := service.UpdateTask(task); err != nil {
		t.Fatalf("UpdateTask() returned error: %v", err)
	}

	if mockRepo.Tasks[0].Title != "更新後タスク" {
		t.Errorf("Expected Title '更新後タスク', got %q", mockRepo.Tasks[0].Title)
	}
}

func TestUpdateTask(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	task := &model.Task{ID: 1, Title: "削除タスク", Status: "pending"}
	mockRepo.Tasks = append(mockRepo.Tasks, *task)
	service := NewTaskService(mockRepo)

	if err := service.DeleteTask(task.ID); err != nil {
		t.Fatalf("DeleteTask() returned error: %v", err)
	}

	if len(mockRepo.Tasks) != 0 {
		t.Errorf("Expected 0 tasks after delete, got %d", len(mockRepo.Tasks))
	}
}
