package user

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

func (r *repo) Register(user *User) (*User, error) {
	tx := r.DB.Begin()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrDatabase
	}
	tx.Commit()
	return user, nil
}

func (r *repo) FindByID(id uint32) (*User, error) {
	tx := r.DB.Begin()
	user := &User{}
	if err := tx.Where("id=?", id).Find(user).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrNotFound
	}
	tx.Commit()
	return user, nil
}

func (r *repo) FindByEmail(email string) (*User, error) {
	tx := r.DB.Begin()
	user := &User{}
	if err := tx.Where("email=?", email).Find(user).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, pkg.ErrNotFound
	}
	tx.Commit()
	return user, nil
}
func (r *repo) DoesEmailExist(email string) bool {
	user := &User{}
	//RecordNotFound returns true, if email does not exist,we reverse the boolean before returning
	return !r.DB.Where("email=?", email).Find(user).RecordNotFound()
}
