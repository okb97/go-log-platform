package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/okb97/go-log-platform/internal/repository"
)

type MockLogRepo struct {
	SavedData []byte
	Err       error
}

func NewLogService(repo repository.LogRepository) *LogService {
	return &LogService{repo: repo}
}

func (m *MockLogRepo) SaveLogs(data []byte) error {
	if m.Err != nil {
		return m.Err
	}
	m.SavedData = data
	return nil
}

func TestFetchLogs_N001(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test log"))
	}))
	defer ts.Close()

	mockRepo := &MockLogRepo{}
	logService := NewLogService(mockRepo)

	err := logService.FetchLogs(ts.URL, "dummy.log")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if string(mockRepo.SavedData) != "test log" {
		t.Errorf("expected saved data 'test log', got %s", string(mockRepo.SavedData))
	}
}
