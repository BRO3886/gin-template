package article

//Service is another abstraction over repository
type Service interface {
	CreateArticle(article *Article) (*Article, error)
	GetUserArticles(userID uint32) (*[]Article, error)
	GetAllArticles() (*[]Article, error)
	GetArticleByID(ID uint32) (*Article, error)
	DeleteArticle(ID uint32) (bool, error)
}

type service struct {
	repo Repository
}

//NewService impl
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func (s service) CreateArticle(article *Article) (*Article, error) {
	return s.repo.CreateArticle(article)
}

func (s service) GetUserArticles(userID uint32) (*[]Article, error) {
	return s.repo.GetUserArticles(userID)
}

func (s service) GetAllArticles() (*[]Article, error) {
	return s.repo.GetAllArticles()
}

func (s service) GetArticleByID(ID uint32) (*Article, error) {
	return s.repo.GetArticlebyID(ID)
}

func (s service) DeleteArticle(ID uint32) (bool, error) {
	return s.repo.DeleteArticle(ID)
}
