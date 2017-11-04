package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/non-player-games/metro-simulation/store"
	"github.com/non-player-games/metro-simulation/ticker"
	"github.com/rcliao/redux"
	"github.com/rcliao/sql-unit-test/web"
)

func init() {
	store.Init()
}

func main() {
	duration := 5 * time.Second
	// TODO: replace duration to 1*time.Minute
	ticker := ticker.NewTicker(duration, simulationTick(store.Store))
	ticker.Run()

	r := mux.NewRouter()
	r.HandleFunc("/hello", web.Hello()).Methods("GET")

	log.Println("Running web server at port 8000")
	http.ListenAndServe(":8000", r)
}

func simulationTick(store *redux.Store) func(t time.Time) error {
	return func(t time.Time) error {
		store.Dispatch(redux.Action{Type: "TRAIN_DEPARTURE"})
		store.Dispatch(redux.Action{Type: "RIDER_SHOWS_UP_STATION"})
		store.Dispatch(redux.Action{Type: "RIDER_DEPARTURE_TRAIN"})
		store.Dispatch(redux.Action{Type: "RIDER_ARRIVAL_TRAIN"})
		return nil
	}
}
