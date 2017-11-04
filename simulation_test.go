package simulation

import "testing"

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
	if newTrainState.CurrentStation.ID != line.Stations[1].ID {
		t.Error("train should departure to the correct next station", newTrainState.CurrentStation)
	}
}
