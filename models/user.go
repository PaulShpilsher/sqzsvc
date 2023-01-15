package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) SaveUser() (*User, error) {
	if err := Database.Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}
