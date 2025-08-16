package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/okb97/go-log-platform/internal/repository"
)

func FetchLogs(serverURL string, logPath string) error {
	client := &http.Client{}

	url := fmt.Sprintf("%s/%s", serverURL, logPath)

	response, err := client.Get(url)
	if err != nil {

	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {

	}

	body, err := io.ReadAll(response.Body)
	if err != nil {

	}

	if err := repository.SaveLogs(body); err != nil {

	}

	return nil
}
