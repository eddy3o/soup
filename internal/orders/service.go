package orders

type Service interface {
	CreateOrder(userID string, req OrderCreateRequest) (Order, error)
	GetOrdersByUserID(userID string) ([]Order, error)
	GetOrderDetailsByOrderID(orderID string) ([]OderDetail, error)
	PutOrderStatusCancelled(orderID string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateOrder(userID string, req OrderCreateRequest) (Order, error) {
	return s.repo.CreateOrder(userID, req)
}

func (s *service) GetOrdersByUserID(userID string) ([]Order, error) {
	return s.repo.GetOrdersByUserID(userID)
}

func (s *service) GetOrderDetailsByOrderID(orderID string) ([]OderDetail, error) {
	return s.repo.GetOrderDetailsByOrderID(orderID)
}

func (s *service) PutOrderStatusCancelled(orderID string) error {
	return s.repo.PutOrderStatusCancelled(orderID)
}
