package structs

type PokemonListResponse struct {
	PokemonList []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}