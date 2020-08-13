package article

import (
	"github.com/BRO3886/gin-learn/pkg"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func (r *repo) CreateArticle(article *Article) (*Article, error) {
	tx := r.DB.Begin()

	if err := tx.Create(article).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrDatabase
	}
	tx.Commit()
	return article, nil
}

func (r *repo) GetUserArticles(userID uint32) (*[]Article, error) {
	var articles []Article

	tx := r.DB.Begin()

	if err := tx.Where("user_id=?", userID).Find(&articles).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrDatabase
	}
	tx.Commit()
	return &articles, nil
}

func (r *repo) GetArticlebyID(ID uint32) (*Article, error) {
	var article Article

	tx := r.DB.Begin()
	if err := tx.Where("id=?", ID).Find(&article).Error; err != nil {
		return nil, pkg.ErrDatabase
	}
	tx.Commit()
	return &article, nil
}
