package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskRouter(t *testing.T) {
	router := TaskRouter()

	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expecter status 200, got %d", w.Code)
	}
}
