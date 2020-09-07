package model

// PlanetAPI : PlanetAPI Object
type PlanetAPI struct {
	Nome    string   `json:"name"`
	Filmes  []string `json:"films"`
	Clima   string   `json:climate`
	Terreno string   `json:terrain`
}
