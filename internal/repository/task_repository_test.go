package repository

import (
	"testing"
	"time"

	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/model"
)

func TestGetAllTasks(t *testing.T) {
	testDB := db.InitTestDB()
	db.DB = testDB

	tasks := []model.Task{
		{Title: "買い物に行く", Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Title: "Go学習", Status: "completed", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.DB.Create(&tasks).Error; err != nil {
		t.Fatalf("Failed to seed test data: %v", err)
	}

	var got []model.Task
	if err := testDB.Find(&got).Error; err != nil {
		t.Fatalf("Failed to fetch tasks: %v", err)
	}

	got, err := GetAllTasks()
	if err != nil {
		t.Fatalf("GetAllTasks() returned error: %v", err)
	}

	//t.Log("Fetched tasks:", got)

	if len(got) != len(tasks) {
		t.Errorf("Expected %d tasks, but got %d", len(tasks), len(got))
	}

	for i, task := range tasks {
		if got[i].Title != task.Title {
			t.Errorf("Expected title %q but got %q", task.Title, got[i].Title)
		}
		if got[i].Status != task.Status {
			t.Errorf("Expected status %q but got %q", task.Status, got[i].Status)
		}
	}

}
