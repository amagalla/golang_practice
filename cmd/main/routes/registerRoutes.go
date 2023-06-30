package routes

import (
	"golang_practice/cmd/main/routes/physicians"
	"golang_practice/cmd/main/routes/pokemon"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/pokemon", pokemon.GetPokemonList)

	router.POST("/insertPhysicians", physicians.InsertPhysicians)
}