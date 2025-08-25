package repository

import (
	"testing"

	"github.com/okb97/go-log-platform/internal/model"
)

func TestGetAllTasks(t *testing.T) {
	mockRepo := NewMockTaskRepository()
	mockRepo.Tasks = []model.Task{
		{ID: 1, Title: "買い物", Status: "pending"},
		{ID: 2, Title: "Go学習", Status: "completed"},
	}

	got, err := mockRepo.GetAllTasks()
	if err != nil {
		t.Fatalf("GetAllTasks() returned error: %v", err)
	}

	if len(got) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(got))
	}

	if got[0].Title != "買い物" || got[1].Title != "Go学習" {
		t.Errorf("Unexpected task titles: %+v", got)
	}
}

func TestCreateTask(t *testing.T) {
	mockRepo := NewMockTaskRepository()
	task1 := &model.Task{Title: "買い物", Status: "pending"}
	task2 := &model.Task{Title: "Go学習", Status: "completed"}

	if err := mockRepo.CreateTask(task1); err != nil {
		t.Fatalf("CreateTask() returned error: %v", err)
	}
	if err := mockRepo.CreateTask(task2); err != nil {
		t.Fatalf("CreateTask() returned error: %v", err)
	}

	if len(mockRepo.Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(mockRepo.Tasks))
	}
}

func TestDeleteTask(t *testing.T) {
	mockRepo := NewMockTaskRepository()

	task1 := &model.Task{Title: "買い物", Status: "pending"}
	task2 := &model.Task{Title: "Go学習", Status: "completed"}

	// タスク作成
	_ = mockRepo.CreateTask(task1)
	_ = mockRepo.CreateTask(task2)

	// task1 を削除
	if err := mockRepo.DeleteTask(task1.ID); err != nil {
		t.Fatalf("DeleteTask() returned error: %v", err)
	}

	if len(mockRepo.Tasks) != 1 {
		t.Errorf("Expected 1 task after delete, got %d", len(mockRepo.Tasks))
	}

	if mockRepo.Tasks[0].ID != task2.ID {
		t.Errorf("Remaining task ID mismatch, got %d", mockRepo.Tasks[0].ID)
	}
}

func TestUpdateTask(t *testing.T) {
	mockRepo := NewMockTaskRepository()

	// まずタスクを作成（CreateTask で ID 付与）
	task := &model.Task{Title: "更新前テストタスク", Status: "pending"}
	_ = mockRepo.CreateTask(task)

	// 更新
	task.Title = "更新後テストタスク"
	if err := mockRepo.UpdateTask(task); err != nil {
		t.Fatalf("UpdateTask() returned error: %v", err)
	}

	// 検証
	if len(mockRepo.Tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(mockRepo.Tasks))
	}
	if mockRepo.Tasks[0].Title != "更新後テストタスク" {
		t.Errorf("Expected Title '更新後テストタスク', got %q", mockRepo.Tasks[0].Title)
	}
}
