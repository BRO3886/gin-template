package article

import "github.com/jinzhu/gorm"

//Article corresponds to articles users can write
type Article struct {
	gorm.Model
	ID          uint32 `json:"id" gorm:"primary_key"`
	UserID      uint32 `json:"user_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Content     string `json:"content"`
	DateCreated string `json:"date_created"`
}
