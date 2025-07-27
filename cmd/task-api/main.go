package main

import (
	"fmt"

	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/repository"
)

func main() {
	db.InitDB()
	repository.SeedTestData()
	tasks, err := repository.GetAllTasks()
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		fmt.Printf("%+v\n", task)
	}
}
