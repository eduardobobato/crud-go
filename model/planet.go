package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Planet : Planet Object
type Planet struct {
	ID             primitive.ObjectID `bson:"_id"`
	Nome           string             `bson:"Nome" json:"Nome"`
	Clima          string             `bson:"Clima" json:"Clima"`
	Terreno        string             `bson:"Terreno" json:"Terreno"`
	CountAparicoes int                `bson:"CountAparicoes" json:"CountAparicoes"`
}
