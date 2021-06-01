package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID      string    `gorm:"column:uuid; primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at; autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at; autoUpdateTime"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.UUID = uuid.New().String()
	return nil
}
