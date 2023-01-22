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
	Email    string    `gorm:"size:256;not null;uniqueIndex;<-:create" json:"email"`
	Password string    `gorm:"size:64;not null;" json:"password"`
	UserUrls []UserUrl `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) GetUserById(id uint) (*User, error) {

	if err := db.First(&u, id).Error; err != nil {
		return u, errors.New("User not found")
	}

	return u, nil
}

func (u *User) GetUserByEmail(email string) (*User, bool) {
	tx := db.Limit(1).Where(&User{Email: email}).Find(&u)
	ok := !errors.Is(tx.Error, gorm.ErrRecordNotFound) && tx.RowsAffected == 1
	return u, ok
}

func (u *User) SaveUser() (*User, error) {
	err := db.Create(&u).Error
	return u, err
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
