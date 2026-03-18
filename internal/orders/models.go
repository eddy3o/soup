package orders

import "time"

type Order struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	OrderDatetime  time.Time `json:"order_date"`
	PickupDatetime time.Time `json:"pickup_date" binding:"required"`
	Status         string    `json:"status"`
	GeneralNotes   string    `json:"general_notes" binding:"max=500"`
	CreatedAt      time.Time `json:"created_at"`
}

type OderDetail struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
	UnitPrice float64   `json:"unit_price" binding:"required,gt=0"`
	Notes     string    `json:"notes" binding:"max=500"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderCreateRequest struct {
	PickupDatetime time.Time    `json:"pickup_date" binding:"required"`
	GeneralNotes   string       `json:"general_notes" binding:"max=500"`
	OrderDetails   []OderDetail `json:"order_details" binding:"required,dive"`
}
