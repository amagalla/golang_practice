package main

import (
	"golang_practice/cmd/main/routes/pokemon"
	"golang_practice/pckg/db"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// mysql connection
	db.InitDB()

	router := gin.Default()

	// Set trusted proxy
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Import routes
	router.GET("/pokemon", pokemon.GetPokemonList)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router.Run(":" + port)
}
