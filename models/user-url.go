package models

type UserUrl struct {
	Model
	LongUrl   string `gorm:"size:255;not null;"`
	ShortCode string `gorm:"size:10;not null;uniqueIndex"`
	UserId    uint   `gorm:"not null;"`
	User      User   `gorm:"references:ID"`
}

func (u *UserUrl) SaveUserUrl() (*UserUrl, error) {
	if err := db.Create(&u).Error; err != nil {
		return &UserUrl{}, err
	}
	return u, nil
}
