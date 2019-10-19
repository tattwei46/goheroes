package handler

import (
	"encoding/json"
	"net/http"

	"../customlogger"
	"github.com/gorilla/mux"
)

type Hero struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	logger := customlogger.GetInstance()
	logger.Println("GetHeroes endpoint triggered")
	var heroes []Hero
	heroes = append(heroes, Hero{ID: "1", Name: "Ironman"})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func GetHero(w http.ResponseWriter, r *http.Request) {
	logger := customlogger.GetInstance()
	logger.Println("GetHero endpoint triggered")
	var heroes []Hero
	heroes = append(heroes,
		Hero{ID: "1", Name: "Ironman"},
		Hero{ID: "2", Name: "Hulk"},
		Hero{ID: "3", Name: "Captain America"})
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, hero := range heroes {
		if hero.ID == params["id"] {
			json.NewEncoder(w).Encode(hero)
		}
	}
}

func DeleteHero(w http.ResponseWriter, r *http.Request) {
	logger := customlogger.GetInstance()
	logger.Println("DeleteHero endpoint triggered")
	var heroes []Hero
	heroes = append(heroes,
		Hero{ID: "1", Name: "Ironman"},
		Hero{ID: "2", Name: "Hulk"},
		Hero{ID: "3", Name: "Captain America"})
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, hero := range heroes {
		if hero.ID == params["id"] {
			heroes = append(heroes[:index], heroes[index+1:]...)
			json.NewEncoder(w).Encode(heroes)
		}
	}
}

func UpdateHero(w http.ResponseWriter, r *http.Request) {
	logger := customlogger.GetInstance()
	logger.Println("UpdateHero endpoint triggered")
	var heroes []Hero
	var updatedHeroes []Hero
	heroes = append(heroes,
		Hero{ID: "1", Name: "Ironman"},
		Hero{ID: "2", Name: "Hulk"},
		Hero{ID: "3", Name: "Captain America"})
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, hero := range heroes {
		if hero.ID == params["id"] {
			var hero Hero
			_ = json.NewDecoder(r.Body).Decode(&hero)
			hero.ID = params["id"]
			updatedHeroes = append(heroes[:index], hero)
			updatedHeroes = append(updatedHeroes, heroes[index+1:]...)
			json.NewEncoder(w).Encode(updatedHeroes)
		}
	}
}

func AddHero(w http.ResponseWriter, r *http.Request) {
	logger := customlogger.GetInstance()
	logger.Println("AddHero endpoint triggered")
	w.Header().Set("Content-Type", "application/json")
	var hero Hero
	_ = json.NewDecoder(r.Body).Decode(&hero)
	json.NewEncoder(w).Encode(&hero)
}
