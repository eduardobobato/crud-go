package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eduardobobato/crud-go/model"
	"github.com/eduardobobato/crud-go/service"

	"github.com/gorilla/mux"
)

var planetSerive = service.PlanetService{}

// respondWithError : Set a error message in response
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

// respondWithJSON : Set message in response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// swagger:route GET /planet planet GetAllPlanets
// Return a list of planets
// responses:
//	200: planetsResponse

// TODO: Incluir optional param de queryParam
// GetAll : Get all planets
func GetAll(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	planets, err := planetSerive.GetAll(params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, planets)
}

// swagger:route GET /planet/{id} planet FindPlanetById
// Return a planet by id
// responses:
//	200: planetResponse

// GetByID : Get planet by id
func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, err := planetSerive.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Planet ID")
		return
	}
	respondWithJSON(w, http.StatusOK, planet)
}

// swagger:route POST /planet planet CreatePlanet
// Create a new planet
//
// responses:
//	200: planetResponse

// Create : Create a planet
func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet model.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	newPlanet, err := planetSerive.Create(planet)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, newPlanet)
}

// swagger:route PUT /planet planet UpdatePlanetById
// Update a planet details by id
//
// responses:
//	200: planetResponse

// Update : Update a planet
func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet model.Planet
	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := planetSerive.Update(params["id"], planet); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": planet.Nome + " atualizado com sucesso!"})
}

// swagger:route DELETE /planet/{id} planet DeletePlanetById
// Delete a planet by id
//
// responses:
//	201: planetResponse

// Delete : Delete a planet
func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := planetSerive.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
