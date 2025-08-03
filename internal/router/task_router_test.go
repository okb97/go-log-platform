package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/model"
)

func TestGetAllTasksRouter(t *testing.T) {
	testDB := db.InitTestDB()
	db.DB = testDB

	tasks := []model.Task{
		{Title: "掃除", Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.DB.Create(&tasks).Error; err != nil {
		t.Fatalf("failed to seed test data: %v", err)
	}

	router := TaskRouter()

	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expecter status 200, got %d", w.Code)
	}
}

func TestCreateTaskRouter(t *testing.T) {
	testDB := db.InitTestDB()
	db.DB = testDB

	router := TaskRouter()

	jsonStr := `{"title":"テストタスク","status":"pending"}`
	req, err := http.NewRequest(http.MethodPost, "/api/task", strings.NewReader(jsonStr))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
	}

	body := w.Body.String()

	if !strings.Contains(body, "テストタスク") {
		t.Errorf("Response body does not contain expected task title, got: %s", body)
	}

}

func TestDeleteTaskRouter(t *testing.T) {
	testDB := db.InitTestDB()
	db.DB = testDB

	task := model.Task{
		Title: "散歩", Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
	if err := db.DB.Create(&task).Error; err != nil {
		t.Fatalf("failed to seed test data: %v", err)
	}

	router := TaskRouter()

	url := fmt.Sprintf("/api/tasks/%d", task.ID)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("DELETE /api/tasks/:id: expected 204 No Content, got %d", w.Code)
	}

}
