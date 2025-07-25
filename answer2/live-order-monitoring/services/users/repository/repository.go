package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"live-order-monitoring/services/users/model"
)

type RepositoryPort interface {
	CreateUser(ctx context.Context, user model.User) error
	FindByEmail(ctx context.Context, email string) (model.User, error)
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) CreateUser(ctx context.Context, user model.User) error {
	_, err := r.db.ExecContext(ctx, `
        INSERT INTO users (id, email, password_hash, role_id, created_at)
        VALUES ($1, $2, $3, $4, now())
    `, user.ID, user.Email, user.PasswordHash, user.RoleID)

	if err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}

	return nil
}

func (r *repositoryAdapter) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRowContext(ctx, `
        SELECT id, email, password_hash, role_id
        FROM users
        WHERE email = $1
    `, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.RoleID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("find user failed: %w", err)
	}

	return user, nil
}
