package models

import "time"

type Transition struct {
	ShortCode     string    `gorm:"type:varchar(16);not null;index"` // TODO: KF to URLEntry
	ClientAddress string    `gorm:"type:varchar(64);not null"`
	CreatedAt     time.Time `gorm:"not null"`
}

func (u *Transition) Save() error {
	return db.Create(&u).Error
}
