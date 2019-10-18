package app

import (
	"fmt"
	"log"
	"net/http"

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
	a.Router.HandleFunc("/", handler)
}
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}
