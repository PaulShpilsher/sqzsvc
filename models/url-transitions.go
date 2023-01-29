package models

import "time"

type UrlTransition struct {
	ShortCode     string    `gorm:"type:varchar(16);not null;index"`
	ClientAddress string    `gorm:"type:varchar(64);not null"`
	CreatedAt     time.Time `gorm:"not null"`
}
