package models

import "time"

// Order represents an order in the system
type Order struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	TotalPrice float64   `json:"total_price" db:"total_price"`
	Status     string    `json:"status" db:"status"` // e.g., "pending", "preparing", "ready", "completed", "cancelled"
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// OrderItem represents a specific item within an order with its quantity at the time of purchase
type OrderItem struct {
	ID         int     `json:"id" db:"id"`
	OrderID    int     `json:"order_id" db:"order_id"`
	MenuItemID int     `json:"menu_item_id" db:"menu_item_id"`
	Quantity   int     `json:"quantity" db:"quantity"`
	Price      float64 `json:"price" db:"price"` // Фиксируем цену на момент заказа
}

// OrderItemInput is used for creating/updating order items from frontend input, where we only need MenuItemID and Quantity. Price will be calculated on the backend based on the current price of the menu item.
type OrderItemInput struct {
	MenuItemID int `json:"menu_item_id" binding:"required" example:"1"`
	Quantity   int `json:"quantity" binding:"required,gt=0" example:"2"`
}

// CreateOrderInput payload to create a new order, containing a list of items with their menu item IDs and quantities. The backend will calculate the total price based on current menu item prices and the quantities ordered.
type CreateOrderInput struct {
	Items []OrderItemInput `json:"items" binding:"required,dive"`
}

// UpdateOrderStatusInput to update order status for baristas and admins to manage order progress. Status should be one of the predefined values like "pending", "preparing", "ready", "completed", or "cancelled".
type UpdateOrderStatusInput struct {
	Status string `json:"status" binding:"required" example:"preparing"`
}
