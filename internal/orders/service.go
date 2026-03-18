package orders

type Service interface {
	CreateOrder(userID string, req OrderCreateRequest) (Order, error)
	GetOrdersByUserID(userID string) ([]Order, error)
	GetOrderDetailsByOrderID(orderID string) ([]OderDetail, error)
	PutOrderStatusCancelled(orderID string) error
}

type service struct {
	repository repository
}

func NewService(repository repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateOrder(userID string, req OrderCreateRequest) (Order, error) {
	return s.repository.CreateOrder(userID, req)
}

func (s *service) GetOrdersByUserID(userID string) ([]Order, error) {
	return s.repository.GetOrdersByUserID(userID)
}

func (s *service) GetOrderDetailsByOrderID(orderID string) ([]OderDetail, error) {
	return s.repository.GetOrderDetailsByOrderID(orderID)
}

func (s *service) PutOrderStatusCancelled(orderID string) error {
	return s.repository.PutOrderStatusCancelled(orderID)
}
