package main

import (
	postgresql "live-order-monitoring/services/orders/db"
	"live-order-monitoring/services/users/server"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("services/users/config/.env")
	if err != nil {
		log.Println("Warning: .env not loaded (might be running in production)")
	}
	// fmt.Println("POSTGRES_HOST =", os.Getenv("POSTGRES_HOST"))

	db, err := postgresql.Postgres()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		os.Exit(1)
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// Router setup
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	server.SetupUserRoutes(router, db)

	err = router.Run(":8889")
	if err != nil {
		panic(err.Error())
	}
}
