package service

import (
	"errors"
	"net/url"

	"github.com/eduardobobato/crud-go/dao"
	"github.com/eduardobobato/crud-go/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlanetService : struct PlanetService
type PlanetService struct{}

var api = SwAPI{}
var planetDao = dao.PlanetDAO{}

// GetAll : Get all planets
func (m PlanetService) GetAll(params url.Values) ([]model.Planet, error) {
	return planetDao.GetAll(params)
}

// GetByID : Get planet by id
func (m PlanetService) GetByID(id string) (model.Planet, error) {
	return planetDao.GetByID(id)
}

// Create : Create a planet
func (m PlanetService) Create(planet model.Planet) (model.Planet, error) {
	planet.ID = primitive.NewObjectID()
	planetAPI := api.FindPlannet(planet.Nome)
	if planetAPI.Filmes != nil {
		planet.CountAparicoes = len(planetAPI.Filmes)
	}
	return planetDao.Create(planet)
}

// Update : Update a planet
func (m PlanetService) Update(id string, planet model.Planet) error {
	return planetDao.Update(id, planet)
}

// Delete : Delete a planet
func (m PlanetService) Delete(id string) error {
	return planetDao.Delete(id)
}

func (m PlanetService) ValidatePlanet(planet *model.Planet) error {
	if planet == nil {
		return errors.New("The planet is empty")
	}
	if planet.Nome == "" {
		return errors.New("The planet name is empty")
	}
	return nil
}
