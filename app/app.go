package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"./handler"
	"github.com/gorilla/mux"
)

type App struct {
	Router     *mux.Router
	FileServer http.Handler
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
	spa := spaHandler{staticPath: "static", indexPath: "index.html"}
	a.Router.PathPrefix("/").Handler(spa)
}

func (a *App) setRouters() {
	a.Get("/hello", a.handleRequest(checkAuthorized(handler.SayHello)))
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

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
