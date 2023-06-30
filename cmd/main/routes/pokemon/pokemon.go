package pokemon

import (
	"encoding/json"
	"golang_practice/cmd/main/structs"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPokemonList(c *gin.Context) {

	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		page = 1
	}

	page_size, err := strconv.Atoi(c.Query("page_size"))

	if err != nil {
		page_size = 1
	}

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

	var pokemonListResponse structs.PokemonListResponse
	err = json.NewDecoder(resp.Body).Decode(&pokemonListResponse)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	totalPages := int(math.Ceil(float64(len(pokemonListResponse.PokemonList)) / float64(page_size)))

	startIndex := (page - 1) * page_size
	endIndex := startIndex + page_size

	if startIndex >= len(pokemonListResponse.PokemonList) {
		c.JSON(http.StatusOK, gin.H{
			"pokemonList": []structs.Pokemon{},
		})

		return
	}

	if endIndex > len(pokemonListResponse.PokemonList) {
		endIndex = len(pokemonListResponse.PokemonList)
	}

	pokemonPage := pokemonListResponse.PokemonList[startIndex:endIndex]

	c.JSON(http.StatusOK, gin.H{
		"pokemonList": pokemonPage,
		"totalPages":  totalPages,
	})
}
