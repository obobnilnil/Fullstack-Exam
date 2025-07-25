package service

import (
	"context"
	"errors"
	"fmt"
	"live-order-monitoring/services/users/dto"
	"live-order-monitoring/services/users/model"
	"live-order-monitoring/services/users/repository"
	"live-order-monitoring/services/users/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServicePort interface {
	Register(ctx context.Context, req dto.RegisterRequest) (model.User, error)
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

type serviceAdapter struct {
	r repository.RepositoryPort
}

func NewServiceAdapter(r repository.RepositoryPort) ServicePort {
	return &serviceAdapter{r: r}
}

func (s *serviceAdapter) Register(ctx context.Context, req dto.RegisterRequest) (model.User, error) {
	if req.Email == "" || req.Password == "" || (req.RoleID != 1 && req.RoleID != 2) {
		return model.User{}, errors.New("email, password, and valid role_id are required")
	}

	_, err := s.r.FindByEmail(ctx, req.Email)
	if err == nil {
		return model.User{}, errors.New("email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := model.User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		PasswordHash: string(hashed),
		RoleID:       req.RoleID,
	}

	if err := s.r.CreateUser(ctx, user); err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *serviceAdapter) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := s.r.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
