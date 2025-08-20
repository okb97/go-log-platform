package model

import "time"

type Log struct {
	ID        uint   `gorm:primaryKey`
	Data      []byte `gorm:"type:blob"`
	CreatedAt time.Time
}
