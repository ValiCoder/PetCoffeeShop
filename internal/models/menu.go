package models

import "time"

// Category Represents item groups like Coffee, Tea, Desserts
type Category struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// MenuItem Represents a product available in the coffee shop
type MenuItem struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	Volume      string    `json:"volume" db:"volume"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	IsAvailable bool      `json:"is_available" db:"is_available"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CreateMenuItemInput Admin payload required to add a new menu item
type CreateMenuItemInput struct {
	Name        string  `json:"name" binding:"required,min=2,max=150"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gte=0"`
	Volume      string  `json:"volume"`
	ImageURL    string  `json:"image_url"`
	CategoryID  int     `json:"category_id" binding:"required"`
}

// ErrorResponse используется для возврата ошибок в формате JSON
type ErrorResponse struct {
	Error string `json:"error"`
}
