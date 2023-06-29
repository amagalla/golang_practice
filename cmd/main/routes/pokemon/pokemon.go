package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PokemonListResponse struct {
	PokemonList []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func GetPokemonList(c *gin.Context) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/")

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Error": "Internal Error",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.AbortWithStatusJSON(400, gin.H{
			"Error": "Failed to retrieve Pokemon",
		})
		return
	}

	var pokemonListResponse PokemonListResponse
	err = json.NewDecoder(resp.Body).Decode(&pokemonListResponse)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pokemonList": pokemonListResponse.PokemonList,
	})
}
