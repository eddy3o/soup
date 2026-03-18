package orders

import "time"

type Order struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	OrderDatetime  time.Time `json:"order_date"`
	PickupDatetime time.Time `json:"pickup_date"`
	Status         string    `json:"status"`
	GeneralNotes   string    `json:"general_notes"`
	CreatedAt      time.Time `json:"created_at"`
}

type OderDetail struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderCreateRequest struct {
	PickupDatetime time.Time    `json:"pickup_date" binding:"required"`
	GeneralNotes   string       `json:"general_notes" binding:"max=500"`
	OrderDetails   []OderDetail `json:"order_details" binding:"required,dive"`
}
