// package planetrouter

// import (
// 	planetDAO "config/dao"
// 	"model"
// 	"testing"

// 	"bou.ke/monkey"
// )

// func TestCreate(t *testing.T) {
// 	dao := planetDAO.PlanetDAO{}
// 	guard := monkey.PatchInstanceMethod(dao.Create, func(planet model.Planet) (model.Planet, error) {
// 		return planet, nil
// 	})
// 	planet := model.Planet{Nome: "1"}
// 	newPlanet, _ := dao.Create(planet)
// 	if newPlanet.Nome != "1" {
// 		t.Errorf("Soma esperada: %d", 4)
// 	}
// 	guard.Unpatch()
// }
