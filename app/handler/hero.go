package handler

import (
	"net/http"

	"../customlogger"
	"github.com/gin-gonic/gin"
)

type Hero struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetHeroes(c *gin.Context) {
	logger := customlogger.GetInstance()
	logger.Println("GetHeroes endpoint triggered")
	var heroes []Hero
	heroes = append(heroes, Hero{ID: "1", Name: "Ironman"})
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, heroes)
}

func GetHero(c *gin.Context) {
	logger := customlogger.GetInstance()
	logger.Println("GetHero endpoint triggered")
	var heroes []Hero
	heroes = append(heroes,
		Hero{ID: "1", Name: "Ironman"},
		Hero{ID: "2", Name: "Hulk"},
		Hero{ID: "3", Name: "Captain America"})
	for _, hero := range heroes {
		if hero.ID == c.Param("ID") {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, hero)
		}
	}
}

func DeleteHero(c *gin.Context) {
	logger := customlogger.GetInstance()
	logger.Println("DeleteHero endpoint triggered")
	var heroes []Hero
	heroes = append(heroes,
		Hero{ID: "1", Name: "Ironman"},
		Hero{ID: "2", Name: "Hulk"},
		Hero{ID: "3", Name: "Captain America"})
	for index, hero := range heroes {
		if hero.ID == c.Param("ID") {
			heroes = append(heroes[:index], heroes[index+1:]...)
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, heroes)
		}
	}
}

func UpdateHero(c *gin.Context) {
	logger := customlogger.GetInstance()
	logger.Println("UpdateHero endpoint triggered")
	var heroes []Hero
	var updatedHeroes []Hero
	heroes = append(heroes,
		Hero{ID: "1", Name: "Ironman"},
		Hero{ID: "2", Name: "Hulk"},
		Hero{ID: "3", Name: "Captain America"})
	for index, hero := range heroes {
		if hero.ID == c.Param("ID") {
			var hero Hero
			c.Bind(&hero)
			hero.ID = c.Param("ID")
			updatedHeroes = append(heroes[:index], hero)
			updatedHeroes = append(updatedHeroes, heroes[index+1:]...)
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, updatedHeroes)
		}
	}
}

func AddHero(c *gin.Context) {
	logger := customlogger.GetInstance()
	logger.Println("AddHero endpoint triggered")
	c.Header("Content-Type", "application/json")
	var hero Hero
	c.Bind(&hero)
	c.JSON(http.StatusOK, hero)
}
