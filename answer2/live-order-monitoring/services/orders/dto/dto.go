package dto

type CreateOrderRequest struct {
	Customer string             `json:"customer" binding:"required"`
	Items    []OrderItemRequest `json:"items" binding:"required,dive,required"`
}

type OrderItemRequest struct {
	ProductName string `json:"product_name" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required,gt=0"`
}

type UpdateOrderRequest struct {
	Status string `json:"status" binding:"required"`
}

type AssignOrderRequest struct {
	StaffID string `json:"staff_id" binding:"required,uuid"`
}
