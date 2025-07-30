package service

import (
	"fmt"
	"testing"

	"github.com/okb97/go-log-platform/db"
)

func TestTaskService(t *testing.T) {

	db.InitDB()
	db.SeedTestData()

	results, err := GetAllTasks()
	if err != nil {
		t.Errorf("%d", err)
	}
	for _, result := range results {
		fmt.Printf("%+v\n", result)
	}

}
