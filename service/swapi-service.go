package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	model "github.com/eduardobobato/crud-go/model"
)

// SwAPIService : interface
type SwAPIService interface {
	FindPlannet(nomePlaneta string) (model.PlanetAPI, error)
	Find(url string) (*model.ReturnAPI, error)
}

// api : Sctruct
type swAPI struct{}

// URL Reference
const (
	APIURL = "https://swapi.dev/api"
	PLANET = "/planets"
)

// NewSwAPIService : return a new SwAPIService
func NewSwAPIService() SwAPIService {
	return &swAPI{}
}

// FindPlannet : Find plannet by Name
func (m *swAPI) FindPlannet(nomePlaneta string) (model.PlanetAPI, error) {
	url := APIURL + PLANET
	var planeta model.PlanetAPI
	hasMath := false
	for !hasMath && url != "" && nomePlaneta != "" {
		var response, err = m.Find(url)
		if err != nil {
			return planeta, err
		}
		for _, value := range response.Planetas {
			if value.Nome == nomePlaneta {
				planeta = value
				hasMath = true
				break
			}
		}
		url = response.URLProximo
	}
	return planeta, nil
}

func (*swAPI) Find(url string) (*model.ReturnAPI, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response model.ReturnAPI
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
