package models

import (
	"sqzsvc/utils"

	"gorm.io/gorm"
)

type UserUrl struct {
	Model
	LongUrl   string `gorm:"size:255;not null;"`
	ShortCode string `gorm:"size:11;not null;uniqueIndex"`
	UserId    uint   `gorm:"not null;"`
	User      User   `gorm:"references:ID"`
}

func (u *UserUrl) Save() (*UserUrl, error) {

	if err := db.Create(&u).Error; err != nil {
		return &UserUrl{}, err
	} else {
		return u, nil
	}
}

func (u *UserUrl) BeforeCreate(tx *gorm.DB) error {

	if code, err := nextInSequence(); err != nil {
		return err
	} else {
		u.ShortCode = utils.NumberToShortCode(code)
		return nil
	}
}
