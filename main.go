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
	"github.com/eduardobobato/crud-go/config"
	"github.com/eduardobobato/crud-go/controller"
	"github.com/eduardobobato/crud-go/dao"
	router "github.com/eduardobobato/crud-go/http"
	"github.com/eduardobobato/crud-go/service"
)

var (
	configAPI        config.Config               = config.NewConfig()
	swAPI            service.SwAPI               = service.NewSwAPI()
	planetDAO        dao.PlanetDao               = dao.NewMongoDAO(configAPI)
	planetService    service.PlanetService       = service.NewPlanetService(planetDAO, swAPI)
	httpRouter       router.Router               = router.NewMuxRouter()
	planetController controller.PlanetController = controller.NewPlanetController(planetService)
)

func main() {
	const port string = ":3333"
	httpRouter.GET("/api/v1/planet", planetController.GetAll)
	httpRouter.GET("/api/v1/planet/{id}", planetController.GetByID)
	httpRouter.POST("/api/v1/planet", planetController.Create)
	httpRouter.PUT("/api/v1/planet/{id}", planetController.Update)
	httpRouter.DELETE("/api/v1/planet/{id}", planetController.Delete)
	httpRouter.SWAGGER("/docs", "/swagger.yaml")
	httpRouter.SERVE(port)
}
