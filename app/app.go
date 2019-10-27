package app

import (
	"log"
	"net/http"

	"./handler"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router     *gin.Engine
	FileServer http.Handler
}

func (a *App) Initialize() {
	//.Router = mux.NewRouter()
	a.Router = gin.Default()
	a.setRouters()
	// spa := spaHandler{staticPath: "static", indexPath: "index.html"}
	// a.Router.PathPrefix("/").Handler(spa)
}

func (a *App) setRouters() {
	api := a.Router.Group("/api")
	{
		api.GET("/heroes", handler.GetHeroes)
		api.GET("/hero/:ID", handler.GetHero)
		api.POST("/hero", handler.AddHero)
		api.DELETE("/hero/:ID", handler.DeleteHero)
		api.PATCH("/hero/:ID", handler.UpdateHero)
	}

	a.Router.Use(static.Serve("/", static.LocalFile("./static", true)))
}

// func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("GET")
// }

// func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("POST")
// }

// func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("DELETE")
// }

// func (a *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("PATCH")
// }

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

// func checkAuthorized(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Authorization Successful")
// 		next(w, r)
// 	})
// }
