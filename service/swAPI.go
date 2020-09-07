package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	model "github.com/eduardobobato/crud-go/model"
	"net/http"
)

// SwAPI : Sctruct
type SwAPI struct{}

// URL Reference
const (
	APIURL = "https://swapi.dev/api"
	PLANET = "/planets"
)

// FindPlannet : Find plannet by Name
func (m SwAPI) FindPlannet(nomePlaneta string) model.PlanetAPI {
	url := APIURL + PLANET
	var planeta model.PlanetAPI
	hasMath := false
	for !hasMath && url != "" {
		var response = m.Find(url)
		for _, value := range response.Planetas {
			if value.Nome == nomePlaneta {
				planeta = value
				hasMath = true
				break
			}
		}
		url = response.URLProximo
	}
	return planeta
}

func (m SwAPI) Find(url string) model.ReturnAPI {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response model.ReturnAPI
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
