package entity

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	Title         string
	Author        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedDate time.Time
}
