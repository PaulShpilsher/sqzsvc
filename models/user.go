package models

import (
	"errors"
	"html"
	"sqzsvc/services"
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

	result := db.Limit(1).Where(&User{Email: email}).Find(&u)
	if result.Error != nil {
		return &User{}, result.Error
	}

	//	errors.Is(result.Error, gorm.ErrRecordNotFound)
	// if err := db.Model(User{}).Where("email = ?", email).First(&u).Error; err != nil {
	// 	return &User{}, err
	// }
	return u, nil
}

func (u *User) SaveUser() (*User, error) {
	if err := db.Create(&u).Error; err != nil {
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
