package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eduardobobato/crud-go/errors"
	"github.com/eduardobobato/crud-go/model"
	"github.com/eduardobobato/crud-go/service"

	"github.com/gorilla/mux"
)

// PlanetController : is a interface for planet controller
type PlanetController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

var planetSerive service.PlanetService

var swService service.SwAPIService

// NewPlanetController : return a new planet controller
func NewPlanetController(service service.PlanetService, sw service.SwAPIService) PlanetController {
	planetSerive = service
	swService = sw
	return &controller{}
}

// swagger:route GET /planet planet GetAllPlanets
// Return a list of planets
// consumes:
//
// responses:
//	200: planetsResponse
//  500: serviceErrorResponse

// GetAll : Get all planets
func (*controller) GetAll(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	planets, err := planetSerive.GetAll(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, errors.ServiceError{Message: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, planets)
}

// swagger:route GET /planet/{id} planet FindPlanetById
// Return a planet by id
// responses:
//	200: planetResponse
//  400: serviceErrorResponse

// GetByID : Get planet by id
func (*controller) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	planet, err := planetSerive.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.ServiceError{Message: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, planet)
}

// swagger:route POST /planet planet CreatePlanet
// Create a new planet
//
// responses:
//	201: planetResponse
//  400: serviceErrorResponse
//  500: serviceErrorResponse

// Create : Create a planet
func (*controller) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet model.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, errors.ServiceError{Message: "Invalid request payload"})
		return
	}
	if err := service.ValidatePlanet(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, errors.ServiceError{Message: err.Error()})
		return
	}
	planetAPI, err := swService.FindPlannet(planet.Nome)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, errors.ServiceError{Message: err.Error()})
		return
	}
	if planetAPI.Filmes != nil {
		planet.CountAparicoes = len(planetAPI.Filmes)
	}
	newPlanet, err := planetSerive.Create(planet)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, errors.ServiceError{Message: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusCreated, newPlanet)
}

// swagger:route PUT /planet planet UpdatePlanetById
// Update a planet details by id
//
// responses:
//	200: planetResponse
//  400: serviceErrorResponse

// Update : Update a planet
func (*controller) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet model.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, errors.ServiceError{Message: "Invalid request payload"})
		return
	}
	if err := service.ValidatePlanet(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, errors.ServiceError{Message: err.Error()})
		return
	}
	if err := planetSerive.Update(params["id"], planet); err != nil {
		respondWithError(w, http.StatusBadRequest, errors.ServiceError{Message: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "Successfully updated " + planet.Nome + "!"})
}

// swagger:route DELETE /planet/{id} planet DeletePlanetById
// Delete a planet by id
//
// responses:
//	200: planetResponse
//  400: serviceErrorResponse

// Delete : Delete a planet
func (*controller) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := planetSerive.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusBadRequest, errors.ServiceError{Message: err.Error()})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "Success!"})
}

// respondWithError : Set a error message in response
func respondWithError(w http.ResponseWriter, code int, msg errors.ServiceError) {
	respondWithJSON(w, code, msg)
}

// respondWithJSON : Set message in response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
