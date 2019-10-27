package handler

import (
	"net/http"

	"../customlogger"
	"../model/hero"
	"github.com/gin-gonic/gin"
)

// type Hero struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

func GetHeroes(heroList *hero.Heroes) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := customlogger.GetInstance()
		logger.Println("GetHeroes endpoint triggered")
		heroes := heroList.GetAll()
		// var heroes []Hero
		// heroes = append(heroes, Hero{ID: "1", Name: "Ironman"})
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, heroes)
	}
}

func GetHero(heroList *hero.Heroes) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := customlogger.GetInstance()
		logger.Println("GetHero endpoint triggered")
		heroes := heroList.GetAll()
		// var heroes []Hero
		// heroes = append(heroes,
		// 	Hero{ID: "1", Name: "Ironman"},
		// 	Hero{ID: "2", Name: "Hulk"},
		// 	Hero{ID: "3", Name: "Captain America"})
		for _, hero := range heroes {
			if hero.ID == c.Param("ID") {
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, hero)
			}
		}
	}
}

func DeleteHero(heroList *hero.Heroes) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := customlogger.GetInstance()
		logger.Println("DeleteHero endpoint triggered")
		heroes := heroList.GetAll()
		// var heroes []Hero
		// heroes = append(heroes,
		// 	Hero{ID: "1", Name: "Ironman"},
		// 	Hero{ID: "2", Name: "Hulk"},
		// 	Hero{ID: "3", Name: "Captain America"})
		for index, hero := range heroes {
			if hero.ID == c.Param("ID") {
				heroes = append(heroes[:index], heroes[index+1:]...)
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, heroes)
			}
		}
	}
}

func UpdateHero(heroList *hero.Heroes) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := customlogger.GetInstance()
		logger.Println("UpdateHero endpoint triggered")
		heroes := heroList.GetAll()
		var updatedHeroes = hero.Heroes{}
		// heroes = append(heroes,
		// 	Hero{ID: "1", Name: "Ironman"},
		// 	Hero{ID: "2", Name: "Hulk"},
		// 	Hero{ID: "3", Name: "Captain America"})
		for index, hero := range heroes {
			if hero.ID == c.Param("ID") {
				c.Bind(&hero)
				hero.ID = c.Param("ID")
				updatedHeroes.Heroes = append(heroes[:index], hero)
				updatedHeroes.Heroes = append(updatedHeroes.Heroes, heroes[index+1:]...)
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusOK, updatedHeroes)
			}
		}
	}
}

func AddHero(heroList *hero.Heroes) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := customlogger.GetInstance()
		logger.Println("AddHero endpoint triggered")
		c.Header("Content-Type", "application/json")
		hero := hero.Hero{}
		c.Bind(&hero)
		c.JSON(http.StatusOK, hero)
	}
}
