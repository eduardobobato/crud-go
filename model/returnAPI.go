package model

// ReturnAPI : ReturnAPI Object
type ReturnAPI struct {
	URLProximo  string      `json:"next"`
	URLAnterior string      `json:"previous"`
	Planetas    []PlanetAPI `json:"results"`
}
