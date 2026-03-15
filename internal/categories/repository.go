package categories

import "soup/internal/store"

type Repository interface {
	GetCategories() ([]Category, error)
}

type repository struct {
	db *store.Database
}

func NewRepository(db *store.Database) Repository {
	return &repository{db: db}
}

func (r repository) GetCategories() ([]Category, error) {
	query := `SELECT id, name, description FROM categories`
	rows, err := r.db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}
