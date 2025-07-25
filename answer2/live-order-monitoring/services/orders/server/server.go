package server

import (
	"database/sql"
	"live-order-monitoring/services/orders/handler"
	"live-order-monitoring/services/orders/repository"
	"live-order-monitoring/services/orders/service"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, db *sql.DB) {
	r := repository.NewRepositoryAdapter(db)
	s := service.NewServiceAdapter(r)
	h := handler.NewHandlerAdapter(s)

	api := router.Group("/api/orders")
	{
		api.POST("", h.CreateOrder)           // POST /api/orders
		api.GET("", h.GetOrders)              // GET /api/orders
		api.PUT("/:id", h.UpdateOrder)        // PUT /api/orders/:id
		api.PUT("/:id/assign", h.AssignOrder) // PUT /api/orders/:id/assign
		api.PUT("/:id/cancel", h.CancelOrder) // PUT /api/orders/:id/cancel
	}
}
