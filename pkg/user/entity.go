package user

import "github.com/jinzhu/gorm"

//User model which stores user information
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"UNIQUE, not null"`
	Password string `json:"password" gorm:"not null"`
}
