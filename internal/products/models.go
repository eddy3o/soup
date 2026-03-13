package products

import "database/sql"

type Product struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	PhotoURL    sql.NullString `json:"photo_url"`
	Available   bool           `json:"available"`
	Category    string         `json:"category"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}
