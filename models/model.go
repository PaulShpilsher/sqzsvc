package models

import "time"

type Model struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}
