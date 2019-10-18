package handler

import (
	"encoding/json"
	"net/http"
)

type Hero struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var heroes []Hero

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	heroes = append(heroes, Hero{ID: "1", Name: "Ironman"})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}
