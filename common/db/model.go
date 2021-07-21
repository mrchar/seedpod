package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	UUID      string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.UUID = uuid.New().String()
	return nil
}
