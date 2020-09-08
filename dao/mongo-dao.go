package dao

import (
	"context"
	"errors"
	"log"
	"net/url"
	"time"

	"github.com/eduardobobato/crud-go/config"

	"github.com/eduardobobato/crud-go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// repo : Sctruct
type repo struct{}

var collection string
var db *mongo.Database

// NewMongoDAO create a new MongoDao
func NewMongoDAO(config config.Config) PlanetDao {
	serverURI, database, collec := config.Read()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverURI))
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(database)
	collection = collec
	return &repo{}
}

// GetAll : Get all planets have optional Query param 'Nome'
func (*repo) GetAll(params *url.Values) ([]model.Planet, error) {
	var results []model.Planet
	filter := bson.M{}
	if params != nil {
		for key, value := range *params {
			if key == "Nome" {
				filter[key] = bson.M{"$in": value}
			}
			if key == "Clima" {
				filter[key] = bson.M{"$in": value}
			}
			if key == "Terreno" {
				filter[key] = bson.M{"$in": value}
			}
		}
	}
	cur, err := db.Collection(collection).Find(context.TODO(), filter, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var elem model.Planet
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.TODO())
	return results, err
}

// GetByID : Get planet by ID
func (*repo) GetByID(id string) (*model.Planet, error) {
	objectID, erro := primitive.ObjectIDFromHex(id)
	if erro != nil {
		return nil, errors.New("Invalid ID")
	}
	filter := bson.M{
		"_id": bson.M{
			"$eq": objectID,
		},
	}
	var planet model.Planet
	err := db.Collection(collection).FindOne(context.TODO(), filter).Decode(&planet)
	return &planet, err
}

// Create : Create a planet
func (*repo) Create(planet model.Planet) (model.Planet, error) {
	_, err := db.Collection(collection).InsertOne(context.TODO(), planet)
	if err != nil {
		return planet, err
	}
	return planet, err
}

// Delete : Delete a planet by ID
func (*repo) Delete(id string) error {
	objectID, erro := primitive.ObjectIDFromHex(id)
	if erro != nil {
		return errors.New("Invalid ID")
	}
	filter := bson.M{
		"_id": bson.M{
			"$eq": objectID,
		},
	}
	_, err := db.Collection(collection).DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return err
}

// Update : Update a planet by ID
func (*repo) Update(id string, planet model.Planet) error {
	objectID, erro := primitive.ObjectIDFromHex(id)
	if erro != nil {
		return errors.New("Invalid ID")
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
