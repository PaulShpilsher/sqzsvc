package models

import (
	"errors"
	"sqzsvc/utils"
	"time"

	"gorm.io/gorm"
)

type Url struct {
	ShortCode     string    `gorm:"type:varchar(16);not null;primarykey"`
	LongUrl       string    `gorm:"type:varchar(4000);not null;uniqueIndex"`
	CreatedAt     time.Time `gorm:"not null"`
	ClientAddress string    `gorm:"type:varchar(64);not null;index"`
}

func (u *Url) Save() error {
	return db.Create(&u).Error
}

func (u *Url) BeforeCreate(tx *gorm.DB) error {

	if code, err := nextInSequence(); err != nil {
		return err
	} else {
		u.ShortCode = utils.NumberToShortCode(code)
		return nil
	}
}

func (u *Url) GetByLongUrl(longUrl string) (*Url, bool) {
	tx := db.Limit(1).Where(&Url{LongUrl: longUrl}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}

func (u *Url) GetByShortCode(shortCode string) (*Url, bool) {
	tx := db.Limit(1).Where(&Url{ShortCode: shortCode}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}
