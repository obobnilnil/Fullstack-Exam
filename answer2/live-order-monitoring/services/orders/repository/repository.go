package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"live-order-monitoring/services/orders/model"
	"log"
)

type RepositoryPort interface {
	CreateOrder(ctx context.Context, order model.Order) error
	GetAllOrders(ctx context.Context) ([]model.Order, error)
	GetOrderByID(ctx context.Context, id string) (model.Order, error)
	UpdateOrderStatus(ctx context.Context, id string, status string) error
	AssignOrderStaff(ctx context.Context, id string, staffID string) error
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) CreateOrder(ctx context.Context, order model.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx failed: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
		INSERT INTO orders (id, customer_name, status, staff_id)
		VALUES ($1, $2, $3, $4)
	`, order.ID, order.Customer, order.Status, order.StaffID)
	if err != nil {
		return fmt.Errorf("insert order failed: %w", err)
	}

	for _, item := range order.Items {
		_, err := tx.ExecContext(ctx, `
		INSERT INTO order_items (order_id, product_name, quantity)
		VALUES ($1, $2, $3)
	`, order.ID, item.ProductName, item.Quantity)
		if err != nil {
			return fmt.Errorf("insert order item failed: %w", err)
		}
	}

	return tx.Commit()
}

func (r *repositoryAdapter) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT o.id, o.customer_name, o.status, o.staff_id, o.created_at, o.updated_at
		FROM orders o
	`)
	if err != nil {
		log.Println("query orders failed:", err)
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.ID, &order.Customer, &order.Status, &order.StaffID, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			log.Println("scan order failed:", err)
			return nil, err
		}

		items, err := r.getOrderItems(ctx, order.ID)
		if err != nil {
			log.Println("get items failed:", err)
			return nil, err
		}
		order.Items = items

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *repositoryAdapter) GetOrderByID(ctx context.Context, id string) (model.Order, error) {
	var order model.Order
	err := r.db.QueryRowContext(ctx, `
		SELECT id, customer_name, status, staff_id, created_at, updated_at
		FROM orders WHERE id = $1
	`, id).Scan(&order.ID, &order.Customer, &order.Status, &order.StaffID, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return order, fmt.Errorf("order not found: %w", err)
		}
		return order, fmt.Errorf("get order failed: %w", err)
	}

	items, err := r.getOrderItems(ctx, order.ID)
	if err != nil {
		return order, err
	}
	order.Items = items

	return order, nil
}

func (r *repositoryAdapter) UpdateOrderStatus(ctx context.Context, id string, status string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE orders SET status = $1 WHERE id = $2`, status, id)
	if err != nil {
		return fmt.Errorf("update order status failed: %w", err)
	}
	return nil
}

func (r *repositoryAdapter) AssignOrderStaff(ctx context.Context, id string, staffID string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE orders SET staff_id = $1 WHERE id = $2`, staffID, id)
	if err != nil {
		return fmt.Errorf("assign order staff failed: %w", err)
	}
	return nil
}

func (r *repositoryAdapter) getOrderItems(ctx context.Context, orderID string) ([]model.OrderItem, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, order_id, product_name, quantity
		FROM order_items WHERE order_id = $1
	`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.OrderItem
	for rows.Next() {
		var item model.OrderItem
		err := rows.Scan(&item.ID, &item.OrderID, &item.ProductName, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
