package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(22);not null"`
	Phone    string `gorm:"type:varchar(22);not null;unique"`
	Password string `gorm:"size:255;not null;"`
}
