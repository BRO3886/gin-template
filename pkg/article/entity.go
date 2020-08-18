package article

import "github.com/jinzhu/gorm"

//Article corresponds to articles users can write
type Article struct {
	gorm.Model
	UserID  uint32 `json:"user_id" gorm:"not null"`
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content"`
}
