package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/okb97/go-log-platform/internal/repository"
)

//test4_3.ymlの確認2

type LogServiceInterface interface {
	FetchLogs(serveURL string, logPath string) error
}

type LogService struct {
	repo repository.LogRepository
}

func (s *LogService) FetchLogs(serverURL string, logPath string) error {
	client := &http.Client{}

	url := fmt.Sprintf("%s/%s", serverURL, logPath)

	response, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch logs from %s: %w", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d from %s", response.StatusCode, url)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if err := s.repo.SaveLogs(body); err != nil {
		return fmt.Errorf("failed to save logs: %w", err)
	}

	return nil
}
