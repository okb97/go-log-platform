package db

import (
	"log"

	"github.com/okb97/go-log-platform/internal/model"
)

func SeedTestData() {
	tasks := []model.Task{
		{Title: "買い物に行く", Status: "pending"},
		{Title: "Go学習", Status: "completed"},
		{Title: "掃除する", Status: "pending"},
	}

	result := DB.Create(&tasks) // 複数件まとめて挿入
	if result.Error != nil {
		log.Fatal("テストデータ挿入失敗:", result.Error)
	}
}
