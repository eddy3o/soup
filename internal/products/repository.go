package products

import "soup/internal/store"

type Repository interface {
	FindAll() ([]Product, error)
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

func (r *repository) FindAll() ([]Product, error) {
	// TODO: Implement database query to fetch all products
	return nil, nil
}

func (r *repository) FindByID(id string) (*Product, error) {
	// TODO: Implement database query to fetch a product by id
	return nil, nil
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
