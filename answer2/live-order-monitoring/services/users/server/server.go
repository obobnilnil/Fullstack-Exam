package server

import (
	"database/sql"
	"live-order-monitoring/services/users/handler"
	"live-order-monitoring/services/users/repository"
	"live-order-monitoring/services/users/service"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, db *sql.DB) {
	r := repository.NewRepositoryAdapter(db)
	s := service.NewServiceAdapter(r)
	h := handler.NewHandlerAdapter(s)

	api := router.Group("/api/users")
	{
		api.POST("/register", h.Register) // POST /api/users/register
		api.POST("/login", h.Login)       // POST /api/users/login
	}
}
