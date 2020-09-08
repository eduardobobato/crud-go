package service

import (
	"errors"
	"net/url"

	"github.com/eduardobobato/crud-go/dao"
	"github.com/eduardobobato/crud-go/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlanetService : interface to service
type PlanetService interface {
	GetAll(params url.Values) ([]model.Planet, error)
	GetByID(id string) (model.Planet, error)
	Create(planet model.Planet) (model.Planet, error)
	Update(id string, planet model.Planet) error
	Delete(id string) error
	ValidatePlanet(planet *model.Planet) error
}

// PlanetService : struct PlanetService
type service struct{}

var api SwAPI
var planetDao dao.PlanetDao

// NewPlanetService : Instance a new service
func NewPlanetService(planetDAO dao.PlanetDao, sw SwAPI) PlanetService {
	api = sw
	planetDao = planetDAO
	return &service{}
}

// GetAll : Get all planets
func (*service) GetAll(params url.Values) ([]model.Planet, error) {
	return planetDao.GetAll(params)
}

// GetByID : Get planet by id
func (*service) GetByID(id string) (model.Planet, error) {
	return planetDao.GetByID(id)
}

// Create : Create a planet
func (*service) Create(planet model.Planet) (model.Planet, error) {
	planet.ID = primitive.NewObjectID()
	planetAPI := api.FindPlannet(planet.Nome)
	if planetAPI.Filmes != nil {
		planet.CountAparicoes = len(planetAPI.Filmes)
	}
	return planetDao.Create(planet)
}

// Update : Update a planet
func (*service) Update(id string, planet model.Planet) error {
	return planetDao.Update(id, planet)
}

// Delete : Delete a planet
func (*service) Delete(id string) error {
	return planetDao.Delete(id)
}

// ValidatePlanet : Validate fields
func (*service) ValidatePlanet(planet *model.Planet) error {
	if planet == nil {
		return errors.New("The planet is empty")
	}
	if planet.Nome == "" {
		return errors.New("The planet name is empty")
	}
	return nil
}
