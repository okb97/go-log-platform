package main

import (
	"github.com/okb97/go-log-platform/db"
	"github.com/okb97/go-log-platform/internal/repository"
	"github.com/okb97/go-log-platform/internal/router"
)

func main() {
	db.InitDB()
	repository.SeedTestData()

	r := router.TaskRouter()

	r.Run(":8080")
}
