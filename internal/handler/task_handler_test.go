package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/model"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/tasks", GetAllTasksHandler)
	r.POST("/task", CreateTaskHandler)

	return r
}

func TestGetAllTasksHandler(t *testing.T) {
	testDB := db.InitTestDB()
	db.DB = testDB

	tasks := []model.Task{
		{Title: "散歩", Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Title: "料理", Status: "completed", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.DB.Create(&tasks).Error; err != nil {
		t.Fatalf("failed to seed test data: %v", err)
	}

	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	body, _ := io.ReadAll(w.Body)
	bodyStr := string(body)
	if len(bodyStr) == 0 {
		t.Error("expected non-empty body")
	}

	if !strings.Contains(bodyStr, "散歩") || !strings.Contains(bodyStr, "料理") {
		t.Errorf("response body missing expected task titles, got: %s", bodyStr)
	}
}

func TestCreateTaskHandler(t *testing.T) {
	testDB := db.InitTestDB()
	db.DB = testDB

	router := setupRouter()

	jsonStr := `{"title":"テストタスク","status":"pending"}`

	req, err := http.NewRequest(http.MethodPost, "/task", strings.NewReader(jsonStr))
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
