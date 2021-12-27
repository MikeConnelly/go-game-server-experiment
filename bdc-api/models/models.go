package models

import (
	"time"

	"gorm.io/gorm"
)

func GenerateISOString() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z07:00")
}

// BeforeCreate will set Base struct before every insert
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	t := GenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t
	return nil
}

// AfterUpdate will update the Base struct after every update
func (base *Base) AfterUpdate(tx *gorm.DB) error {
	base.UpdatedAt = GenerateISOString()
	return nil
}

// Base contains common columns for all tables
type Base struct {
	ID        uint   `gorm:"unique;primaryKey;autoIncrement"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
