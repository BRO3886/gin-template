package article

//Repository interface for a layer of abstraction over the db functions
type Repository interface {
	CreateArticle(article *Article) (*Article, error)
	GetUserArticles(userID uint32) (*[]Article, error)
	GetAllArticles() (*[]Article, error)
	GetArticlebyID(ID uint32) (*Article, error)
	DeleteArticle(ID uint32) (bool, error)
}
