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
	"fmt"
	"log"
	"net/http"

	"github.com/eduardobobato/crud-go/config"
	"github.com/eduardobobato/crud-go/controller"
	"github.com/eduardobobato/crud-go/dao"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gorilla/mux"
)

var planetDao = dao.PlanetDAO{}
var configAPI = config.Config{}

func init() {
	configAPI.Read()
	planetDao.ServerURI = configAPI.ServerURI
	planetDao.Database = configAPI.Database
	planetDao.Collection = configAPI.Collection
	planetDao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/planet", controller.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/planet/{id}", controller.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/planet", controller.Create).Methods("POST")
	r.HandleFunc("/api/v1/planet/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/api/v1/planet/{id}", controller.Delete).Methods("DELETE")

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	r.Handle("/docs", sh)
	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	var port = ":3333"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
