package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"src/projector/db"
	"src/projector/handlers"
	"src/projector/middleware"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	router := gin.Default()

	router.POST("/keys", handlers.CreateKey)

	auth := router.Group("/")
	auth.Use(middleware.APIKeyAuth())
	{
		auth.POST("/orders", handlers.CreateOrder)
		auth.GET("/orders", handlers.GetOrders)
		auth.PUT("/orders/:id", handlers.UpdateOrder)
		auth.DELETE("/orders/:id", handlers.DeleteOrder)
	}

	router.Run("localhost:8080")
}
