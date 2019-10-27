package testhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../handler"
	"../model/hero"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {

	heroList := hero.New()
	heroList.Add(hero.Hero{ID: "1", Name: "Ironman"})
	heroList.Add(hero.Hero{ID: "2", Name: "Captain America"})
	heroList.Add(hero.Hero{ID: "3", Name: "Hulk"})

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/heroes", handler.GetHeroes(heroList))
		api.GET("/hero/:ID", handler.GetHero(heroList))
		api.POST("/hero", handler.AddHero(heroList))
		api.DELETE("/hero/:ID", handler.DeleteHero(heroList))
		api.PATCH("/hero/:ID", handler.UpdateHero(heroList))
	}

	return router
}

func TestGetHeroes(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/heroes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	expected := "[{\"id\":\"1\",\"name\":\"Ironman\"},{\"id\":\"2\",\"name\":\"Captain America\"},{\"id\":\"3\",\"name\":\"Hulk\"}]\n"
	actual := w.Body.String()
	assert.Equal(t, expected, actual)

}
