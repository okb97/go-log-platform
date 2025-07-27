package service

import (
	"fmt"
	"testing"

	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/repository"
)

func TestTaskService(t *testing.T) {

	db.InitDB()
	repository.SeedTestData()

	results, err := GetAllTasks()
	if err != nil {
		t.Errorf("%d", err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

}
