package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/internal/model"
	"github.com/okb97/go-log-platform/internal/service"
)

type MockTaskService struct {
	Tasks []model.Task
}

func (m *MockTaskService) GetAllTasks() ([]model.Task, error) {
	return m.Tasks, nil
}

func (m *MockTaskService) CreateTask(task *model.Task) error {
	task.ID = uint(len(m.Tasks) + 1)
	m.Tasks = append(m.Tasks, *task)
	return nil
}

func (m *MockTaskService) DeleteTask(id uint) error {
	for i, t := range m.Tasks {
		if t.ID == id {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockTaskService) UpdateTask(task *model.Task) error {
	for i, t := range m.Tasks {
		if t.ID == task.ID {
			m.Tasks[i] = *task
			return nil
		}
	}
	return nil
}
func setupRouterWithService(taskService service.TaskServiceInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	handler := NewTaskHandler(taskService)

	r.GET("/api/tasks", handler.GetAllTasksHandler)
	r.POST("/api/task", handler.CreateTaskHandler)
	r.DELETE("/api/tasks/:id", handler.DeleteTaskHandler)
	r.PUT("/api/tasks/:id", handler.UpdateTaskHandler)

	return r
}

func TestGetAllTasksHandler(t *testing.T) {
	mockService := &MockTaskService{
		Tasks: []model.Task{
			{ID: 1, Title: "散歩", Status: "pending"},
			{ID: 2, Title: "料理", Status: "completed"},
		},
	}
	router := setupRouterWithService(mockService)

	req, _ := http.NewRequest(http.MethodGet, "/api/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	body, _ := io.ReadAll(w.Body)
	bodyStr := string(body)
	if !strings.Contains(bodyStr, "散歩") || !strings.Contains(bodyStr, "料理") {
		t.Errorf("response body missing expected task titles, got: %s", bodyStr)
	}
}

func TestCreateTaskHandler(t *testing.T) {
	mockService := &MockTaskService{}
	router := setupRouterWithService(mockService)

	jsonStr := `{"title":"テストタスク","status":"pending"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/task", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
	}

	if len(mockService.Tasks) != 1 || mockService.Tasks[0].Title != "テストタスク" {
		t.Errorf("Task was not created correctly in service")
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	mockService := &MockTaskService{
		Tasks: []model.Task{
			{ID: 1, Title: "散歩", Status: "pending"},
		},
	}
	router := setupRouterWithService(mockService)

	req, _ := http.NewRequest(http.MethodDelete, "/api/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status %d but got %d", http.StatusNoContent, w.Code)
	}

	if len(mockService.Tasks) != 0 {
		t.Errorf("Task was not deleted in service")
	}
}

func TestUpdateTaskHandler(t *testing.T) {
	mockService := &MockTaskService{
		Tasks: []model.Task{
			{ID: 1, Title: "散歩", Status: "pending"},
		},
	}
	router := setupRouterWithService(mockService)

	jsonStr := `{"title":"更新タスク","status":"completed"}`
	req, _ := http.NewRequest(http.MethodPut, "/api/tasks/1", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status %d but got %d", http.StatusNoContent, w.Code)
	}

	if mockService.Tasks[0].Title != "更新タスク" || mockService.Tasks[0].Status != "completed" {
		t.Errorf("Task was not updated in service")
	}
}
