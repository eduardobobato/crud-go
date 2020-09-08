package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	model "github.com/eduardobobato/crud-go/model"
)

// SwAPI : interface
type SwAPI interface {
	FindPlannet(nomePlaneta string) model.PlanetAPI
	Find(url string) model.ReturnAPI
}

// api : Sctruct
type swAPI struct{}

// URL Reference
const (
	APIURL = "https://swapi.dev/api"
	PLANET = "/planets"
)

// NewSwAPI : return a new NewSwAPI
func NewSwAPI() SwAPI {
	return &swAPI{}
}

// FindPlannet : Find plannet by Name
func (m *swAPI) FindPlannet(nomePlaneta string) model.PlanetAPI {
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

func (*swAPI) Find(url string) model.ReturnAPI {
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
