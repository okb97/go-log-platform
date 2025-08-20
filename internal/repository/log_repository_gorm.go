package repository

import (
	"github.com/okb97/go-log-platform/internal/model"
	"gorm.io/gorm"
)

type GormLogRepository struct {
	db *gorm.DB
}

func NewGormLogRepository(db *gorm.DB) *GormLogRepository {
	return &GormLogRepository{db: db}
}

func (r *GormLogRepository) SaveLogs(data []byte) error {
	log := model.Log{Data: data}
	return r.db.Create(&log).Error
}
