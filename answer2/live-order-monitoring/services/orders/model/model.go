package model

import "time"

type Order struct {
	ID        string      `json:"id"`
	Customer  string      `json:"customer"`
	Status    string      `json:"status"`
	StaffID   *string     `json:"staff_id,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Items     []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
	ID          string `json:"id"`
	OrderID     string `json:"order_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
