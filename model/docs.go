package model

import "github.com/eduardobobato/crud-go/errors"

// PlanetFilter in query
// swagger:parameters GetAllPlanets
type PlanetFilter struct {
	// Planet name.
	//
	// in: query
	Nome string `json:"Nome"`

	// Planet climate.
	//
	// in: query
	Clima string `json:"Clima"`

	// Planet ground.
	//
	// in: query
	Terreno string `json:"Terreno"`
}

// A list of planet
// swagger:response planetsResponse
type planetsResponseWrapper struct {
	// All current planet
	// in: body
	Body []Planet
}

// A single of planet
// swagger:response planetResponse
type planetResponseWrapper struct {
	// A planet
	// in: body
	Body Planet
}

// A error response
// swagger:response serviceErrorResponse
type serviceErrorResponseWrapper struct {
	// A planet
	// in: body
	Body errors.ServiceError
}
