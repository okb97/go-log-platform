package repository

import (
	"log"

	"github.com/okb97/go-log-platform/db"
)

func SeedTestData() {
	query := `
        INSERT INTO tasks (title, status)
        VALUES
        ('買い物に行く', 'pending'),
        ('Go学習', 'completed'),
        ('掃除する', 'pending');
    `

	_, err := db.DB.Exec(query)
	if err != nil {
		log.Fatal("テストデータ挿入失敗:", err)
	}
}
