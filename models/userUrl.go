package models

import (
	"errors"
	"sqzsvc/utils"

	"gorm.io/gorm"
)

type UserUrl struct {
	Model
	UserID    uint   `gorm:"not null"`
	ShortCode string `gorm:"type:varchar(11);not null;uniqueIndex"`
	LongUrl   string `gorm:"type:VARCHAR(4000);not null;"`
}

func (u *UserUrl) Save() error {
	return db.Create(&u).Error
}

func (u *UserUrl) BeforeCreate(tx *gorm.DB) error {

	if code, err := nextInSequence(); err != nil {
		return err
	} else {
		u.ShortCode = utils.NumberToShortCode(code)
		return nil
	}
}

func (u *UserUrl) GetByUserAndUrl() (*UserUrl, bool) {
	tx := db.Limit(1).Where(&UserUrl{UserID: u.UserID, LongUrl: u.LongUrl}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}
