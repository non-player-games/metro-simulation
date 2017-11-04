package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	simulation "github.com/non-player-games/metro-simulation"
	"github.com/non-player-games/metro-simulation/store"
)

// Homepage is DTO sending to homepage
type Homepage struct {
	Trains   []simulation.Train
	Stations []simulation.Station
	Lines    []simulation.Line
}

// Hello says hello
func Hello() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Metro-Simulation!")
	})
}

// CurrentState returns the current state in JSON
func CurrentState() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store.Store.GetState())
	})
}

// Assets serves the static assets (js & css)
func Assets() http.Handler {
	return http.StripPrefix("/assets", http.FileServer(http.Dir("./web/assets")))
}

// Index renders the index page for submitting SQL queries to test
func Index() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./web/templates/index.html")
		if err != nil {
			panic(err)
		}
		state := store.Store.GetState()
		dto := Homepage{
			Trains:   state["trains"].([]simulation.Train),
			Stations: state["stations"].([]simulation.Station),
			Lines:    state["lines"].([]simulation.Line),
		}
		t.Execute(w, dto)
	})
}
