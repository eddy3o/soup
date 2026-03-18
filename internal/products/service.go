package products

type Service interface {
	FindAll(pagination PaginationParams) ([]Product, int, error)
	FindByID(id string) (*Product, error)
	Create(product Product) (*Product, error)
	Update(id string, product Product) (*Product, error)
	Delete(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) FindAll(pagination PaginationParams) ([]Product, int, error) {
	return s.repo.FindAll(pagination)
}

func (s *service) FindByID(id string) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(product Product) (*Product, error) {
	return s.repo.Create(product)
}

func (s *service) Update(id string, product Product) (*Product, error) {
	return s.repo.Update(id, product)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
