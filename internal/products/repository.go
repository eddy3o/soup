package products

import "soup/internal/store"

type Repository interface {
	FindAll(pagination PaginationParams) ([]Product, int, error)
	FindByID(id string) (*Product, error)
	Create(product Product) (*Product, error)
	Update(id string, product Product) (*Product, error)
	Delete(id string) error
}

type repository struct {
	db *store.Database
}

func NewRepository(db *store.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAll(pagination PaginationParams) ([]Product, int, error) {
	query := `
    SELECT 
			products.id, products.name, products.description, products.price, 
			products.photo_url, products.available, categories.name, 
			products.created_at, products.updated_at,
			COUNT(*) OVER() AS total_count
		FROM products
		INNER JOIN categories ON products.category_id = categories.id
    ORDER BY products.id ASC
		LIMIT $1 OFFSET $2
  `
	var products []Product
	var total int
	rows, err := r.db.DB.Query(query, pagination.Limit, (pagination.Page-1)*pagination.Limit)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.PhotoURL,
			&product.Available,
			&product.Category,
			&product.CreatedAt,
			&product.UpdatedAt,
			&total,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *repository) FindByID(id string) (*Product, error) {
	query := `
		SELECT products.id, products.name, products.description, products.price, products.photo_url, products.available, categories.name, products.created_at, products.updated_at
		FROM products
		INNER JOIN categories on products.category_id = categories.id
		WHERE products.id = $1
	`
	var product Product
	err := r.db.DB.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.PhotoURL,
		&product.Available,
		&product.Category,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) Create(product Product) (*Product, error) {
	// TODO: Implement database query to create a new product
	return nil, nil
}

func (r *repository) Update(id string, product Product) (*Product, error) {
	// TODO: Implement database query to update a product by id
	return nil, nil
}

func (r *repository) Delete(id string) error {
	// TODO: Implement database query to delete a product by is_admin
	return nil
}
