package service

import (
	"context"
	"errors"
	"fmt"
	"live-order-monitoring/services/orders/dto"
	"live-order-monitoring/services/orders/model"
	"live-order-monitoring/services/orders/repository"

	"github.com/google/uuid"
)

type ServicePort interface {
	CreateOrder(ctx context.Context, req dto.CreateOrderRequest) (model.Order, error)
	GetAllOrders(ctx context.Context) ([]model.Order, error)
	UpdateOrder(ctx context.Context, id string, req dto.UpdateOrderRequest) (model.Order, error)
	AssignOrder(ctx context.Context, id string, staffID string) (model.Order, error)
	CancelOrder(ctx context.Context, id string) (model.Order, error)
}

type serviceAdapter struct {
	r repository.RepositoryPort
}

func NewServiceAdapter(r repository.RepositoryPort) ServicePort {
	return &serviceAdapter{r: r}
}

func (s *serviceAdapter) CreateOrder(ctx context.Context, req dto.CreateOrderRequest) (model.Order, error) {
	if req.Customer == "" || len(req.Items) == 0 {
		return model.Order{}, errors.New("customer name and items are required")
	}

	orderID := uuid.New().String()
	// now := time.Now()

	var items []model.OrderItem
	for _, i := range req.Items {
		items = append(items, model.OrderItem{
			OrderID:     orderID,
			ProductName: i.ProductName,
			Quantity:    i.Quantity,
		})
	}

	order := model.Order{
		ID:       orderID,
		Customer: req.Customer,
		Status:   "pending",
		// CreatedAt: now,
		// UpdatedAt: now,
		Items: items,
	}

	if err := s.r.CreateOrder(ctx, order); err != nil {
		return model.Order{}, fmt.Errorf("failed to create order: %w", err)
	}
	// _ = PublishOrderUpdate(order.ID)
	PublishOrderCreated(order)
	return order, nil
}

func (s *serviceAdapter) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	order, err := s.r.GetAllOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all orders: %w", err)
	}
	return order, nil
}

func (s *serviceAdapter) UpdateOrder(ctx context.Context, id string, req dto.UpdateOrderRequest) (model.Order, error) {
	order, err := s.r.GetOrderByID(ctx, id)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to get order by ID %s: %w", id, err)
	}

	if err := s.r.UpdateOrderStatus(ctx, id, req.Status); err != nil {
		return model.Order{}, fmt.Errorf("failed to update order status for ID %s: %w", id, err)
	}

	order.Status = req.Status
	// _ = PublishOrderUpdate(id)
	PublishOrderUpdated(order)
	return order, nil
}

func (s *serviceAdapter) AssignOrder(ctx context.Context, id string, staffID string) (model.Order, error) {
	order, err := s.r.GetOrderByID(ctx, id)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to get order by ID %s: %w", id, err)
	}

	if err := s.r.AssignOrderStaff(ctx, id, staffID); err != nil {
		return model.Order{}, fmt.Errorf("failed to assign staff to order ID %s: %w", id, err)
	}

	order.StaffID = &staffID
	// _ = PublishOrderUpdate(id)
	PublishOrderAssigned(id, staffID)
	return order, nil
}

func (s *serviceAdapter) CancelOrder(ctx context.Context, id string) (model.Order, error) {
	order, err := s.r.GetOrderByID(ctx, id)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to get order by ID %s: %w", id, err)
	}

	order.Status = "cancelled"
	// order.UpdatedAt = time.Now()

	if err := s.r.UpdateOrderStatus(ctx, id, "cancelled"); err != nil {
		return model.Order{}, fmt.Errorf("failed to cancel order ID %s: %w", id, err)
	}
	// _ = PublishOrderUpdate(id)
	PublishOrderCancelled(id)
	return order, nil
}
