package service

import (
	"net/url"
	"testing"

	"github.com/eduardobobato/crud-go/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDAO struct {
	mock.Mock
}

func (mock *MockDAO) Create(planet model.Planet) (model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.Planet), args.Error(1)
}
func (mock *MockDAO) GetAll(params *url.Values) ([]model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Planet), args.Error(1)
}

func (mock *MockDAO) GetByID(id string) (model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.Planet), args.Error(1)
}
func (mock *MockDAO) Delete(id string) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockDAO) Update(id string, planet model.Planet) error {
	args := mock.Called()
	return args.Error(0)
}

func TestValidateEmptyPlanet(t *testing.T) {
	testService := NewPlanetService(nil, nil)

	err := testService.ValidatePlanet(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The planet is empty", err.Error())
}

func TestValidateEmptyNamePlanet(t *testing.T) {
	testService := NewPlanetService(nil, nil)
	planet := model.Planet{Nome: "", Terreno: "", Clima: "", CountAparicoes: 0}
	err := testService.ValidatePlanet(&planet)

	assert.NotNil(t, err)
	assert.Equal(t, "The planet name is empty", err.Error())
}

func TestGetAll(t *testing.T) {
	// Setup expectation
	mockDao := new(MockDAO)
	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockDao.On("GetAll").Return([]model.Planet{planet}, nil)

	testService := NewPlanetService(mockDao, nil)

	result, _ := testService.GetAll(nil)

	mockDao.AssertExpectations(t)

	assert.Equal(t, "A", result[0].Nome)
	assert.Equal(t, "B", result[0].Terreno)
	assert.Equal(t, "C", result[0].Clima)
	assert.Equal(t, 0, result[0].CountAparicoes)
}
