package app

import (
	"fmt"
	"log"
	"net/http"

	"./handler"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/", a.handleRequest(checkAuthorized(handler.SayHello)))
	a.Get("/heroes", a.handleRequest(handler.GetHeroes))
	a.Get("/hero/{id}", a.handleRequest(handler.GetHero))
	a.Post("/hero", a.handleRequest(handler.AddHero))
	a.Delete("/hero/{id}", a.handleRequest(handler.DeleteHero))
	a.Patch("/hero/{id}", a.handleRequest(handler.UpdateHero))

}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PATCH")
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (a *App) handleRequest(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func checkAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authorization Successful")
		next(w, r)
	})
}
