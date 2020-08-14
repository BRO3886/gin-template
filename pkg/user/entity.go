package user

import "github.com/jinzhu/gorm"

//User model which stores user information
type User struct {
	gorm.Model
	UserID   uint32 `json:"user_id"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"UNIQUE, not null"`
	Password string `json:"password" gorm:"not null"`
}
