package user

import "github.com/jinzhu/gorm"

//User model which stores user information
type User struct {
	gorm.Model
	ID       uint32 `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"UNIQUE"`
	Password string `json:"password"`
}
