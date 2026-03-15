package categories

type Service interface {
	GetCategories() ([]Category, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetCategories() ([]Category, error) {
	return s.repo.GetCategories()
}
