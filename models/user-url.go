package models

type UserUrl struct {
	Model
	LongUrl   string `gorm:"size:255;not null;"`
	ShortCode string `gorm:"size:10;not null;uniqueIndex"`
	UserId    uint   `gorm:"not null;"`
	User      User   `gorm:"references:ID"`
}
