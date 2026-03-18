package orders

import "soup/internal/store"

type Repository interface {
	CreateOrder(userID string, req OrderCreateRequest) (Order, error)
	GetOrdersByUserID(userID string) ([]Order, error)
	GetOrderDetailsByOrderID(orderID string) ([]OderDetail, error)
	PutOrderStatusCancelled(orderID string) error
}

type repository struct {
	db *store.Database
}

func NewRepository(db *store.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateOrder(userID string, req OrderCreateRequest) (Order, error) {
	// Implement the logic to create an order in the database
	return Order{}, nil
}

func (r *repository) GetOrdersByUserID(userID string) ([]Order, error) {
	// Implement the logic to retrieve orders by user ID from the database
	return []Order{}, nil
}

func (r *repository) GetOrderDetailsByOrderID(orderID string) ([]OderDetail, error) {
	// Implement the logic to retrieve order details by order ID from the database
	return []OderDetail{}, nil
}

func (r *repository) PutOrderStatusCancelled(orderID string) error {
	// Implement the logic to update the order status to cancelled in the database
	return nil
}
