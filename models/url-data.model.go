package models

import (
	"errors"
	"sqzsvc/utils"
	"time"

	"gorm.io/gorm"
)

type UrlData struct {
	ID            uint      `gorm:"primarykey"`
	CreatedAt     time.Time `gorm:"not null"`
	ShortCode     string    `gorm:"type:varchar(11);not null;uniqueIndex"`
	Url           string    `gorm:"type:varchar(4000);not null;unique"`
	ClientAddress string    `gorm:"type:varchar(64);not null;index"`
}

func (u *UrlData) Save() error {
	return db.Create(&u).Error
}

func (u *UrlData) BeforeCreate(tx *gorm.DB) error {

	if code, err := nextInSequence(); err != nil {
		return err
	} else {
		u.ShortCode = utils.NumberToShortCode(code)
		return nil
	}
}

func (u *UrlData) GetByUrl(url string) (*UrlData, bool) {
	tx := db.Limit(1).Where(&UrlData{Url: url}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}

func (u *UrlData) GetByShortCode(shortCode string) (*UrlData, bool) {
	tx := db.Limit(1).Where(&UrlData{ShortCode: shortCode}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}
