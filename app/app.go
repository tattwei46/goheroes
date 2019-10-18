package app

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
}

func (a *App) Run(port string) {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}
