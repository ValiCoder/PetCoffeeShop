package repository

import (
	"coffee-project/internal/models"
	"database/sql"
)

type MenuRepository interface {
	GetAll() ([]models.MenuItem, error)
	Create(item *models.MenuItem) error
}

type postgresMenuRepository struct {
	db *sql.DB
}

func NewPostgresMenuRepository(db *sql.DB) MenuRepository {
	return &postgresMenuRepository{db: db}
}

func (r *postgresMenuRepository) GetAll() ([]models.MenuItem, error) {
	query := `SELECT id, name, description, price, volume, image_url, category_id, is_available, created_at, updated_at 
	          FROM menu_items WHERE is_available = true ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.MenuItem
	for rows.Next() {
		var item models.MenuItem
		err := rows.Scan(
			&item.ID, &item.Name, &item.Description, &item.Price,
			&item.Volume, &item.ImageURL, &item.CategoryID,
			&item.IsAvailable, &item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *postgresMenuRepository) Create(item *models.MenuItem) error {
	query := `INSERT INTO menu_items (name, description, price, volume, image_url, category_id, is_available, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING id, created_at, updated_at`

	return r.db.QueryRow(
		query, item.Name, item.Description, item.Price,
		item.Volume, item.ImageURL, item.CategoryID, item.IsAvailable,
	).Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt)
}
