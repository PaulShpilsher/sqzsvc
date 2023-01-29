package models

import (
	"errors"
	"sqzsvc/utils"
	"time"

	"gorm.io/gorm"
)

type UrlEntry struct {
	ShortCode     string    `gorm:"type:varchar(16);not null;primarykey"`
	Url           string    `gorm:"type:varchar(4000);not null;uniqueIndex"`
	CreatedAt     time.Time `gorm:"not null"`
	ClientAddress string    `gorm:"type:varchar(64);not null;index"`
}

func (u *UrlEntry) Save() error {
	return db.Create(&u).Error
}

func (u *UrlEntry) BeforeCreate(tx *gorm.DB) error {

	if code, err := nextInSequence(); err != nil {
		return err
	} else {
		u.ShortCode = utils.NumberToShortCode(code)
		return nil
	}
}

func (u *UrlEntry) GetByUrl(url string) (*UrlEntry, bool) {
	tx := db.Limit(1).Where(&UrlEntry{Url: url}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}

func (u *UrlEntry) GetByShortCode(shortCode string) (*UrlEntry, bool) {
	tx := db.Limit(1).Where(&UrlEntry{ShortCode: shortCode}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}
