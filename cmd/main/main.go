package main

import (
	"golang_practice/cmd/main/routes"
	"golang_practice/pckg/db"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	router.Use(cors.Default())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
