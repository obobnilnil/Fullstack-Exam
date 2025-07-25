package handler

import (
	"live-order-monitoring/services/users/dto"
	"live-order-monitoring/services/users/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type handlerAdapter struct {
	s service.ServicePort
}

func NewHandlerAdapter(s service.ServicePort) HandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	user, err := h.s.Register(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered", "user": gin.H{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.RoleID,
	}})
}

func (h *handlerAdapter) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	token, err := h.s.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
