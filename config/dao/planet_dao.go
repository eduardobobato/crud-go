package dao

import (
	"context"
	"log"
	model "github.com/eduardobobato/crud-go/model"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PlanetDAO : Sctruct
type PlanetDAO struct {
	ServerURI  string
	Database   string
	Collection string
}

var collection string
var db *mongo.Database

// Connect to Mongo
func (m *PlanetDAO) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.ServerURI))
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(m.Database)
	collection = m.Collection
}

// GetAll : Get all planets have optional Query param 'Nome'
func (m *PlanetDAO) GetAll(params url.Values) ([]model.Planet, error) {
	var results []model.Planet
	filter := bson.M{}
	for key, value := range params {
		if key == "Nome" {
			filter["Nome"] = bson.M{"$in": value}
		}
		if key == "Clima" {
			filter["Clima"] = bson.M{"$in": value}
		}
		if key == "Terreno" {
			filter["Terreno"] = bson.M{"$in": value}
		}
	}
	cur, err := db.Collection(collection).Find(context.TODO(), filter, options.Find())
	if err != nil {
		log.Fatal(err)
	} else {
		for cur.Next(context.TODO()) {
			var elem model.Planet
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}

			results = append(results, elem)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		cur.Close(context.TODO())
	}
	return results, err
}

// GetByID : Get planet by ID
func (m *PlanetDAO) GetByID(id string) (model.Planet, error) {
	objectID, erro := primitive.ObjectIDFromHex(id)
	if erro != nil {
		log.Fatal(erro)
	}
	filter := bson.M{
		"_id": bson.M{
			"$eq": objectID,
		},
	}
	var planet model.Planet
	err := db.Collection(collection).FindOne(context.TODO(), filter).Decode(&planet)
	return planet, err
}

// Create : Create a planet
func (m *PlanetDAO) Create(planet model.Planet) (model.Planet, error) {
	_, err := db.Collection(collection).InsertOne(context.TODO(), planet)
	if err != nil {
		return planet, err
	}
	return planet, err
}

// Delete : Delete a planet by ID
func (m *PlanetDAO) Delete(id string) error {
	objectID, erro := primitive.ObjectIDFromHex(id)
	if erro != nil {
		log.Fatal(erro)
		return erro
	}
	filter := bson.M{
		"_id": bson.M{
			"$eq": objectID,
		},
	}
	_, err := db.Collection(collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Update : Update a planet by ID
func (m *PlanetDAO) Update(id string, planet model.Planet) error {
	objectID, erro := primitive.ObjectIDFromHex(id)
	if erro != nil {
		log.Fatal(erro)
		return erro
	}
	filter := bson.M{
		"_id": bson.M{
			"$eq": objectID,
		},
	}
	update := bson.M{
		"$set": bson.M{
			"Nome":    planet.Nome,
			"Clima":   planet.Clima,
			"Terreno": planet.Terreno,
		},
	}
	_, err := db.Collection(collection).UpdateOne(context.Background(), filter, update)
	return err
}
