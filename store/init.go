package store

import (
	simulation "github.com/non-player-games/metro-simulation"
	"github.com/rcliao/redux"
)

// Store is the central state management
var Store *redux.Store

// Init initializes the default state
func Init() {
	state := make(map[string]interface{})
	stations := map[string]simulation.Station{
		"MAPLE_STATION": simulation.Station{
			ID:       0,
			Location: simulation.Location{X: 1, Y: 1},
			Name:     "Maple Station",
		},
		"Pine Station": simulation.Station{
			Location: simulation.Location{X: 5, Y: 1},
			Name:     "Pine Station",
			ID:       1,
		},
		"Mahogany Station": simulation.Station{
			Location: simulation.Location{X: 8, Y: 2},
			Name:     "Mahogany Station",
			ID:       2,
		},
		"Palm Station": simulation.Station{
			Location: simulation.Location{X: 14, Y: 2},
			Name:     "Palm Station",
			ID:       3,
		},
		"Ash Station": simulation.Station{
			Location: simulation.Location{X: 5, Y: 4},
			Name:     "Ash Station",
			ID:       4,
		},
		"Cedar Station": simulation.Station{
			Location: simulation.Location{X: 14, Y: 5},
			Name:     "Cedar Station",
			ID:       5,
		},
		"Redwood Station": simulation.Station{
			Location: simulation.Location{X: 2, Y: 6},
			Name:     "Redwood Station",
			ID:       6,
		},
		"Elm Station": simulation.Station{
			Location: simulation.Location{X: 8, Y: 4},
			Name:     "Elm Station",
			ID:       7,
		},
		"Holly Station": simulation.Station{
			Location: simulation.Location{X: 12, Y: 6},
			Name:     "Holly Station",
			ID:       8,
		},
		"Oak Station": simulation.Station{
			Location: simulation.Location{X: 5, Y: 5},
			Name:     "Oak Station",
			ID:       9,
		},
	}
	state["stations"] = getStationSlice(stations)
	state["lines"] = []simulation.Line{
		simulation.Line{
			Stations: []simulation.Station{
				stations["MAPLE_STATION"],
				stations["PINE_STATION"],
				stations["MAHOGANY_STATION"],
			},
			Name: "Tomato",
		},
		simulation.Line{
			Stations: []simulation.Station{
				stations["MAHOGANY_STATION"],
				stations["PALM_STATION"],
				stations["CEDAR_STATION"],
			},
			Name: "Avocado",
		},
		simulation.Line{
			Stations: []simulation.Station{
				stations["MAPLE_STATION"],
				stations["REDWOOD_STATION"],
				stations["OAK_STATION"],
				stations["ASH_STATION"],
				stations["PINE_STATION"],
			},
			Name: "Blueberry",
		},
		simulation.Line{
			Stations: []simulation.Station{
				stations["ASH_STATION"],
				stations["ELM_STATION"],
				stations["MAHOGANY_STATION"],
			},
			Name: "Orange",
		},
		simulation.Line{
			Stations: []simulation.Station{stations["ELM_STATION"], stations["HOLLY_STATION"], stations["CEDAR_STATION"]},
			Name:     "Banana",
		},
	}
	// TODO: add reducers later
	reducers := []redux.Reducer{}
	Store = redux.NewStore(state, reducers)
}

func getStationSlice(m map[string]simulation.Station) []simulation.Station {
	result := []simulation.Station{}
	for _, s := range m {
		result = append(result, s)
	}
	return result
}
