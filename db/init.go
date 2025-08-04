package db

import (
	"fmt"
	"log"

	"github.com/okb97/go-log-platform/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbPath = "task.db"

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB接続失敗：", err)
	}

	err = DB.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatal("マイグレーション失敗：", err)
	}

	fmt.Println("SQLite(GORM) 初期化完了")

}

func InitTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test DB: %v", err)
	}

	err = db.AutoMigrate(
		&model.Task{},
	)
	if err != nil {
		log.Fatalf("failed to migrate test DB: %v", err)
	}

	return db

}
