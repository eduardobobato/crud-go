package dao

import (
	"net/url"

	"github.com/eduardobobato/crud-go/model"
)

// PlanetDao : DAO interface
type PlanetDao interface {
	GetAll(params *url.Values) ([]model.Planet, error)
	GetByID(id string) (*model.Planet, error)
	Create(planet model.Planet) (model.Planet, error)
	Delete(id string) error
	Update(id string, planet model.Planet) error
}
