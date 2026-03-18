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
	query := `
	INSERT INTO orders (user_id, order_datetime, pickup_datetime, status, general_notes)
	VALUES ($1, NOW(), $2, 'pending', $3)
	RETURNING id, user_id, order_datetime, pickup_datetime, status, general_notes, created_at;
	`
	var order Order
	err := r.db.DB.QueryRow(query, userID, req.PickupDatetime, req.GeneralNotes).Scan(
		&order.ID,
		&order.UserID,
		&order.OrderDatetime,
		&order.PickupDatetime,
		&order.Status,
		&order.GeneralNotes,
		&order.CreatedAt,
	)
	if err != nil {
		return Order{}, err
	}
	query = `
	INSERT INTO order_details (order_id, product_id, quantity, unit_price, notes)
	VALUES ($1, $2, $3, $4, $5);
	`
	for _, detail := range req.OrderDetails {
		_, err := r.db.DB.Exec(query, order.ID, detail.ProductID, detail.Quantity, detail.UnitPrice, detail.Notes)
		if err != nil {
			return Order{}, err
		}
	}
	return order, nil
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
