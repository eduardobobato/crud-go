package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Planet : Planet Object

// Planet defines the structure for an API planet
// swagger:model
type Planet struct {
	// the id for the planet
	//
	// required: false
	ID primitive.ObjectID `bson:"_id"`

	// the name for this planet
	//
	// required: true
	Nome string `bson:"Nome" json:"Nome"`

	// the climate for this planet
	//
	// required: true
	Clima string `bson:"Clima" json:"Clima"`

	// the ground for this planet
	//
	// required: true
	Terreno string `bson:"Terreno" json:"Terreno"`

	// the number of appearances in swatter filts
	CountAparicoes int `bson:"CountAparicoes" json:"CountAparicoes"`
}
