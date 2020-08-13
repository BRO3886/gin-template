package article

import (
	"github.com/BRO3886/gin-learn/pkg"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

//NewDatabaseRepo adds a layer of abstraction so that everything
//in the interface Repository can be accessed only via the instance of this function
func NewDatabaseRepo(db *gorm.DB) Repository {
	return &repo{DB: db}
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

func (r *repo) GetAllArticles() (*[]Article, error) {
	var articles []Article

	tx := r.DB.Begin()

	if err := tx.Find(&articles).Error; err != nil {
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

func (r *repo) DeleteArticle(ID uint32) (bool, error) {
	var article Article
	tx := r.DB.Begin()
	if err := tx.Where("id=?", ID).Delete(&article).Error; err != nil {
		return false, pkg.ErrDatabase
	}
	tx.Commit()
	return true, nil
}
