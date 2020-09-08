package service

import (
	"net/url"
	"testing"

	"github.com/eduardobobato/crud-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func (mock *MockDAO) GetByID(id string) (*model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.Planet), args.Error(1)
}
func (mock *MockDAO) Delete(id string) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockDAO) Update(id string, planet model.Planet) error {
	args := mock.Called()
	return args.Error(0)
}

func TestGetAll(t *testing.T) {
	// Setup expectation
	mockDao := new(MockDAO)
	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockDao.On("GetAll").Return([]model.Planet{planet}, nil)

	testService := NewPlanetService(mockDao)

	result, _ := testService.GetAll(nil)

	mockDao.AssertExpectations(t)

	assert.Equal(t, "A", result[0].Nome)
	assert.Equal(t, "B", result[0].Terreno)
	assert.Equal(t, "C", result[0].Clima)
	assert.Equal(t, 0, result[0].CountAparicoes)
}

func TestGetAllByName(t *testing.T) {
	// Setup expectation
	mockDao := new(MockDAO)
	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockDao.On("GetAll").Return([]model.Planet{planet}, nil)

	testService := NewPlanetService(mockDao)

	params, _ := url.ParseQuery("Nome=A")
	result, _ := testService.GetAll(&params)

	mockDao.AssertExpectations(t)

	assert.Equal(t, "A", result[0].Nome)
	assert.Equal(t, "B", result[0].Terreno)
	assert.Equal(t, "C", result[0].Clima)
	assert.Equal(t, 0, result[0].CountAparicoes)
}

func TestCreate(t *testing.T) {
	// Setup expectation
	mockDao := new(MockDAO)
	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockDao.On("Create").Return(planet, nil)

	testService := NewPlanetService(mockDao)

	result, err := testService.Create(planet)

	mockDao.AssertExpectations(t)

	assert.Equal(t, "A", result.Nome)
	assert.Equal(t, "B", result.Terreno)
	assert.Equal(t, "C", result.Clima)
	assert.Equal(t, 0, result.CountAparicoes)
	assert.NotNil(t, result.ID)
	assert.Nil(t, err)
}

func TestGetByID(t *testing.T) {
	mockDao := new(MockDAO)
	stringID := "5f554eb79de81bcb733dbdb0"
	objectID, _ := primitive.ObjectIDFromHex(stringID)

	planet := model.Planet{ID: objectID, Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockDao.On("GetByID").Return(&planet, nil)

	testService := NewPlanetService(mockDao)

	result, err := testService.GetByID(stringID)

	mockDao.AssertExpectations(t)

	assert.Equal(t, objectID, result.ID)
	assert.Equal(t, "A", result.Nome)
	assert.Equal(t, "B", result.Terreno)
	assert.Equal(t, "C", result.Clima)
	assert.Equal(t, 0, result.CountAparicoes)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	mockDao := new(MockDAO)
	stringID := "5f554eb79de81bcb733dbdb0"

	mockDao.On("Delete").Return(nil)

	testService := NewPlanetService(mockDao)

	err := testService.Delete(stringID)

	mockDao.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	mockDao := new(MockDAO)

	mockDao.On("Update").Return(nil)

	testService := NewPlanetService(mockDao)

	stringID := "5f554eb79de81bcb733dbdb0"
	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}

	err := testService.Update(stringID, planet)

	mockDao.AssertExpectations(t)

	assert.Nil(t, err)
}
