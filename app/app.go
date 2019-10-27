package app

import (
	"fmt"
	"log"
	"net/http"

	"./handler"
	"./model/hero"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router     *gin.Engine
	FileServer http.Handler
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.setRouters()
}

func (a *App) setRouters() {

	heroList := hero.New()
	heroList.Add(hero.Hero{ID: "1", Name: "Ironman"})
	heroList.Add(hero.Hero{ID: "2", Name: "Captain America"})
	heroList.Add(hero.Hero{ID: "3", Name: "Hulk"})
	fmt.Println(heroList.GetAll())

	api := a.Router.Group("/api")
	api.Use(AuthenticationRequired())
	{
		api.GET("/heroes", handler.GetHeroes(heroList))
		api.GET("/hero/:ID", handler.GetHero(heroList))
		api.POST("/hero", handler.AddHero(heroList))
		api.DELETE("/hero/:ID", handler.DeleteHero(heroList))
		api.PATCH("/hero/:ID", handler.UpdateHero(heroList))
	}

	a.Router.Use(static.Serve("/", static.LocalFile("./static", true)))
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authorization Successful")
		c.Next()
	}
}
