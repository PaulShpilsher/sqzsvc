package models

import (
	"html"
	"sqzsvc/services"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	if err := Database.Model(User{}).Where("email = ?", email).First(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) SaveUser() (*User, error) {
	if err := Database.Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := services.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	//remove spaces in username
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

}
