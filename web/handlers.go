package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/non-player-games/metro-simulation/store"
)

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
