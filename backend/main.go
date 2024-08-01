package main

import (
	"log"
	"github.com/gin-contrib/cors"
	"bookstore/routers"
	"github.com/gin-gonic/gin"


)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow all origins for testing purposes. Replace with specific URL for production.
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type","ngrok-skip-browser-warning"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	routers.InitRoutes(router)
	router.Run("localhost:8000")
	log.Println("Server started at :8000")
	
}

