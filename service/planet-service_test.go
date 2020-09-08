package service

import (
	"net/url"

	"github.com/eduardobobato/crud-go/model"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Create(planet model.Planet) (model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.Planet), args.Error(1)
}
func (mock *MockRepository) GetAll(params url.Values) ([]model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Planet), args.Error(1)
}

// func (mock *MockRepository) GetByID(id string) (model.Planet, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.([]model.Planet), args.Error(1)
// }
// func (mock *MockRepository) Delete(id string) error                      {}
// func (mock *MockRepository) Update(id string, planet model.Planet) error {}

// func TestValidateEmptyPlanet(t *testing.T) {
// 	testService := NewPlanetService(nil, nil)

// 	err := testService.ValidatePlanet(nil)

// 	assert.NotNil(t, err)
// 	assert.Equal(t, "The planet is empty", err.Error())
// }

// func TestValidateEmptyNamePlanet(t *testing.T) {
// 	testService := NewPlanetService(nil, nil)
// 	planet := model.Planet{Nome: "", Terreno: "", Clima: "", CountAparicoes: 0}
// 	err := testService.ValidatePlanet(&planet)

// 	assert.NotNil(t, err)
// 	assert.Equal(t, "The planet name is empty", err.Error())
// }

// func TestGetAll(t *testing.T) {
// 	mockRepo := new(MockRepository)
// 	// Setup expectation
// 	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
// 	mockRepo.On("GetAll").Return([]model.Planet{planet}, nil)

// 	testService := NewPlanetService(mockRepo)

// 	result, _ = testService.GetAll()

// 	mockRepo.AssertExpectations(t)

// 	assert.Equal(t, "A", result[0].Nome)
// 	assert.Equal(t, "B", result[0].Terreno)
// 	assert.Equal(t, "C", result[0].Clima)
// 	assert.Equal(t, 0, result[0].CountAparicoes)
// }
