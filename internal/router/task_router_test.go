package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/model"
)

func TestTaskRouter(t *testing.T) {
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
