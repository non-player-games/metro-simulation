package simulation

import (
	"math/rand"
	"testing"
)

func TestSimulation(t *testing.T) {
	rand.Seed(1337)
	// for seed 1337, rand.Intn(100) will return 38 on first
	events := []Event{
		Event{
			Chance: 20,
			Value:  "something1",
		},
		Event{
			Chance: 20,
			Value:  "this thing",
		},
		Event{
			Chance: 60,
			Value:  "should not happen",
		},
	}
	if EventSimulation(events).(string) != "this thing" {
		t.Error("event simulation error, should return second item (40 mark)")
	}
}

func TestTrain(t *testing.T) {
	line := Line{
		Stations: []Station{
			Station{
				ID:       0,
				Location: Location{X: 1, Y: 1},
				Name:     "Maple Station",
				Riders:   []Rider{},
			},
			Station{
				Location: Location{X: 5, Y: 1},
				Name:     "Pine Station",
				ID:       1,
				Riders:   []Rider{},
			},
			Station{
				Location: Location{X: 8, Y: 2},
				Name:     "Mahogany Station",
				ID:       2,
				Riders:   []Rider{},
			},
		},
		Name: "Tomato",
	}
	train := Train{
		CurrentStation: line.Stations[0],
		Line:           line,
		Riders:         []Rider{},
		Direction:      true,
	}
	newTrainState := train.Departure()
	if newTrainState.CurrentStation.ID != line.Stations[1].ID || !newTrainState.Direction {
		t.Error("train should departure to the correct next station", newTrainState)
	}
	reverseTrainState := train.Departure().Departure().Departure()
	if reverseTrainState.CurrentStation.ID != line.Stations[1].ID || reverseTrainState.Direction {
		t.Error("train should reach the end and reverse back", reverseTrainState)
	}
}
