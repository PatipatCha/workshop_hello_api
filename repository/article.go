package repository

type Article struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

type ArticleRepository interface {
	GetAll() ([]Article, error)
	GetById(int) (*Article, error)
}
