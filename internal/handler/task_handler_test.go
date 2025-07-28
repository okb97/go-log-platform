package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/repository"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/tasks", GetAllTasksHandler)

	return r
}

func TestGetAllTasksHandler(t *testing.T) {
	db.InitDB()
	repository.SeedTestData()

	router := setupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	body, _ := ioutil.ReadAll(w.Body)
	if len(body) == 0 {
		t.Error("expected non-empty body")
	}
}
