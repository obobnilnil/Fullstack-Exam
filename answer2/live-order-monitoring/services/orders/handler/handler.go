package handler

import (
	"live-order-monitoring/services/orders/dto"
	"live-order-monitoring/services/orders/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandlerPort interface {
	CreateOrder(c *gin.Context)
	GetOrders(c *gin.Context)
	UpdateOrder(c *gin.Context)
	AssignOrder(c *gin.Context)
	CancelOrder(c *gin.Context)
}

type handlerAdapter struct {
	s service.ServicePort
}

func NewHandlerAdapter(s service.ServicePort) HandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	order, err := h.s.CreateOrder(c.Request.Context(), req)
	if err != nil {
		log.Printf("CreateOrder failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create order", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "data": order})
}

func (h *handlerAdapter) GetOrders(c *gin.Context) {
	orders, err := h.s.GetAllOrders(c.Request.Context())
	if err != nil {
		log.Printf("GetOrders failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch orders", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders fetched successfully", "data": orders})
}

func (h *handlerAdapter) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	var req dto.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	order, err := h.s.UpdateOrder(c.Request.Context(), id, req)
	if err != nil {
		log.Printf("UpdateOrder failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update order", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully", "data": order})
}

func (h *handlerAdapter) AssignOrder(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	var req dto.AssignOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	order, err := h.s.AssignOrder(c.Request.Context(), id, req.StaffID)
	if err != nil {
		log.Printf("AssignOrder failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not assign order", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order assigned successfully", "data": order})
}

func (h *handlerAdapter) CancelOrder(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	order, err := h.s.CancelOrder(c.Request.Context(), id)
	if err != nil {
		log.Printf("CancelOrder failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel order", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully", "order": order})
}
