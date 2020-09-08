package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPlannet(t *testing.T) {
	swAPIService := NewSwAPIService()

	resp, _ := swAPIService.FindPlannet("AAA")

	assert.NotNil(t, resp)
	assert.Nil(t, resp.Filmes)
	assert.Empty(t, resp.Nome)
	assert.Empty(t, resp.Terreno)
	assert.Empty(t, resp.Clima)
}

func TestFindPlannetEmptyName(t *testing.T) {
	swAPIService := NewSwAPIService()

	resp, _ := swAPIService.FindPlannet("")

	assert.NotNil(t, resp)
	assert.Nil(t, resp.Filmes)
	assert.Empty(t, resp.Nome)
	assert.Empty(t, resp.Terreno)
	assert.Empty(t, resp.Clima)
}

func TestFindPlannetValidName(t *testing.T) {
	swAPIService := NewSwAPIService()

	resp, _ := swAPIService.FindPlannet("Hoth")

	assert.NotNil(t, resp)
	assert.NotNil(t, resp.Filmes)
	assert.NotEmpty(t, resp.Nome)
}
