package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "task.db"

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("DB接続失敗：", err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        status TEXT NOT NULL DEFAULT 'pending',
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `

	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("テーブル作成失敗:", err)
	}

	fmt.Println("SQLite 初期化完了")

}
