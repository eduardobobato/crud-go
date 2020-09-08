package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/eduardobobato/crud-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) Create(planet model.Planet) (model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.Planet), args.Error(1)
}
func (mock *MockService) GetAll(params *url.Values) ([]model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Planet), args.Error(1)
}

func (mock *MockService) GetByID(id string) (*model.Planet, error) {
	args := mock.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.Planet), args.Error(1)
}
func (mock *MockService) Delete(id string) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockService) Update(id string, planet model.Planet) error {
	args := mock.Called()
	return args.Error(0)
}

type MockSwService struct {
	mock.Mock
}

func (mock *MockSwService) FindPlannet(nomePlaneta string) (model.PlanetAPI, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(model.PlanetAPI), args.Error(1)
}

func (mock *MockSwService) Find(url string) (*model.ReturnAPI, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*model.ReturnAPI), args.Error(1)
}

func TestGetAll(t *testing.T) {
	mockService := new(MockService)
	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockService.On("GetAll").Return([]model.Planet{planet}, nil)

	req, err := http.NewRequest("GET", "/api/v1/planet", nil)
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.GetAll(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAllByName(t *testing.T) {
	mockService := new(MockService)
	planet := model.Planet{Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockService.On("GetAll").Return([]model.Planet{planet}, nil)

	req, err := http.NewRequest("GET", "/api/v1/planet?Nome=A", nil)
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.GetAll(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetByID(t *testing.T) {
	stringID := "5f554eb79de81bcb733dbdb0"
	objectID, _ := primitive.ObjectIDFromHex(stringID)

	mockService := new(MockService)
	planet := model.Planet{ID: objectID, Nome: "A", Terreno: "B", Clima: "C", CountAparicoes: 0}
	mockService.On("GetByID").Return(&planet, nil)

	req, err := http.NewRequest("GET", "/api/v1/planet/"+stringID, nil)
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.GetByID(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetByIDWithInvalidID(t *testing.T) {
	stringID := "a"
	mockService := new(MockService)
	mockService.On("GetByID").Return(nil, errors.New("Invalid ID"))

	req, err := http.NewRequest("GET", "/api/v1/planet/"+stringID, nil)
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.GetByID(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreate(t *testing.T) {
	stringID := "5f554eb79de81bcb733dbdb0"
	objectID, _ := primitive.ObjectIDFromHex(stringID)

	newPlanet := model.Planet{ID: objectID, Nome: "Kamino", Clima: "Frio", Terreno: "Arenoso", CountAparicoes: 1}
	mockService := new(MockService)
	mockService.On("Create").Return(newPlanet, nil)

	planetAPI := model.PlanetAPI{Nome: "Kamino", Terreno: "", Clima: "", Filmes: []string{"http://swapi.dev/api/films/5/"}}
	mockSwService := new(MockSwService)
	mockSwService.On("FindPlannet").Return(planetAPI, nil)

	planet := model.Planet{Nome: "Kamino", Clima: "Frio", Terreno: "Arenoso"}
	requestByte, _ := json.Marshal(planet)

	req, err := http.NewRequest("POST", "/api/v1/planet", bytes.NewReader(requestByte))
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, mockSwService)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestCreateWithoutName(t *testing.T) {

	mockService := new(MockService)
	mockService.On("Create").Return(nil, errors.New("The planet name is empty"))

	planetAPI := model.PlanetAPI{Nome: "", Terreno: "", Clima: "", Filmes: nil}
	mockSwService := new(MockSwService)
	mockSwService.On("FindPlannet").Return(planetAPI, nil)

	planet := model.Planet{Nome: "", Clima: "Frio", Terreno: "Arenoso"}
	requestByte, _ := json.Marshal(planet)

	req, err := http.NewRequest("POST", "/api/v1/planet", bytes.NewReader(requestByte))
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, mockSwService)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateWithoutBody(t *testing.T) {

	mockService := new(MockService)
	mockService.On("Create").Return(nil, errors.New("The planet is empty"))

	planetAPI := model.PlanetAPI{Nome: "", Terreno: "", Clima: "", Filmes: nil}
	mockSwService := new(MockSwService)
	mockSwService.On("FindPlannet").Return(planetAPI, nil)

	requestByte, _ := json.Marshal("")

	req, err := http.NewRequest("POST", "/api/v1/planet", bytes.NewReader(requestByte))
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, mockSwService)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdate(t *testing.T) {
	stringID := "5f554eb79de81bcb733dbdb0"

	mockService := new(MockService)
	mockService.On("Update").Return(nil)

	planet := model.Planet{Nome: "Kamino", Clima: "Frio", Terreno: "Arenoso"}
	requestByte, _ := json.Marshal(planet)

	req, err := http.NewRequest("PUT", "/api/v1/planet/"+stringID, bytes.NewReader(requestByte))
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Update(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestUpdateWithInvalidID(t *testing.T) {
	stringID := "a"

	mockService := new(MockService)
	mockService.On("Update").Return(errors.New("Invalid ID"))

	planet := model.Planet{Nome: "Kamino", Clima: "Frio", Terreno: "Arenoso"}
	requestByte, _ := json.Marshal(planet)

	req, err := http.NewRequest("PUT", "/api/v1/planet/"+stringID, bytes.NewReader(requestByte))
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateWithoutBody(t *testing.T) {
	stringID := "5f554eb79de81bcb733dbdb0"

	mockService := new(MockService)
	mockService.On("Update").Return(errors.New("The planet is empty"))

	requestByte, _ := json.Marshal("")

	req, err := http.NewRequest("PUT", "/api/v1/planet/"+stringID, bytes.NewReader(requestByte))
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestUpdateWithoutName(t *testing.T) {
	stringID := "5f554eb79de81bcb733dbdb0"

	mockService := new(MockService)
	mockService.On("Update").Return(errors.New("The planet name is empty"))

	planet := model.Planet{Nome: "", Clima: "Frio", Terreno: "Arenoso"}
	requestByte, _ := json.Marshal(planet)

	req, err := http.NewRequest("PUT", "/api/v1/planet/"+stringID, bytes.NewReader(requestByte))
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Update(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestDelete(t *testing.T) {
	stringID := "5f554eb79de81bcb733dbdb0"

	mockService := new(MockService)
	mockService.On("Delete").Return(nil)

	req, err := http.NewRequest("DELETE", "/api/v1/planet/"+stringID, nil)
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Delete(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteWithInvalidID(t *testing.T) {
	stringID := "a"
	mockService := new(MockService)
	mockService.On("Delete").Return(errors.New("Invalid ID"))

	req, err := http.NewRequest("DELETE", "/api/v1/planet/"+stringID, nil)
	if err != nil {
		t.Fatal(err)
	}

	testController := NewPlanetController(mockService, nil)

	rr := httptest.NewRecorder()
	ctx := req.Context()
	req = req.WithContext(ctx)

	testController.Delete(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
