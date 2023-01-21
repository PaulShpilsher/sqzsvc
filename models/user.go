package models

import (
	"errors"
	"html"
	"sqzsvc/utils"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	Model
	Email    string `gorm:"size:64;not null;uniqueIndex;<-:create" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u *User) GetUserById(id uint) (*User, error) {

	if err := db.First(&u, id).Error; err != nil {
		return u, errors.New("User not found")
	}

	return u, nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	if err := db.Limit(1).Where(&User{Email: email}).Find(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) SaveUser() (*User, error) {
	if err := db.Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	//turn password into hash
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	//remove spaces in username
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return
}

// func (u *User) AfterFind(tx *gorm.DB) (err error) {
// 	u.Password = ""
// 	return
// }
