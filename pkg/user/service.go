package user

import (
	"fmt"
	"regexp"

	"github.com/BRO3886/gin-learn/pkg"
	"golang.org/x/crypto/bcrypt"
)

//Service interface for abstraction over higher functions
type Service interface {
	Register(user *User) (*User, error)
	Login(email, password string) (*User, error)
	GetUserByID(id uint32) (*User, error)
	GetRepo() Repository
}

type service struct {
	repo Repository
}

//NewService function to get new instance of server
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func (s *service) Register(user *User) (*User, error) {
	allok, err := user.validate()
	if !allok {
		return nil, err
	}

	doesEmailExist := s.repo.DoesEmailExist(user.Email)
	if doesEmailExist {
		return nil, pkg.ErrEmailExists
	}

	user.Password, err = hashpassword(user.Password)
	if err != nil {
		return nil, err
	}

	return s.repo.Register(user)
}

func (s *service) Login(email, password string) (*User, error) {
	user := &User{}
	user, err := s.repo.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	passHash, err := hashpassword(password)
	if err != nil {
		fmt.Println("pass hash error")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetUserByID(id uint32) (*User, error) {
	return s.repo.FindByID(id)
}

func (user *User) validate() (bool, error) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(user.Email) {
		return false, pkg.ErrEmail
	}

	if len(user.Password) < 6 {
		return false, pkg.ErrPassword
	}

	return true, nil
}

func hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
