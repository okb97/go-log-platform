package model

import "time"

type Task struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id`
	Title     string    `json:"title"`
	Status    string    `json:"status"` // 例: "pending", "done"
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
