//  Package CRUD Planet API.
//
// Documentation for Planet API
//
//  Schemes: http
//  BasePath: /api/v1
//  Version: 1.0.0
//  Contact: Eduardo Bobato<eduardobobato@hotmail.com.br>
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//  swagger:meta
package main

import (
	Config "github.com/eduardobobato/crud-go/config"
	PlanetDAO "github.com/eduardobobato/crud-go/config/dao"
	"fmt"
	"log"
	"net/http"
	planetrouter "github.com/eduardobobato/crud-go/router"

	"github.com/gorilla/mux"
)

var dao = PlanetDAO.PlanetDAO{}
var config = Config.Config{}

func init() {
	config.Read()
	dao.ServerURI = config.ServerURI
	dao.Database = config.Database
	dao.Collection = config.Collection
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/planet", planetrouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/planet/{id}", planetrouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/planet", planetrouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/planet/{id}", planetrouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/planet/{id}", planetrouter.Delete).Methods("DELETE")

	var port = ":3333"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
