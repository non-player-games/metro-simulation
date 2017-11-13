package store

import (
	"time"

	simulation "github.com/non-player-games/metro-simulation"
	"github.com/rcliao/redux"
)

// Store is the central state management
var Store *redux.Store

// Init initializes the default state
func Init(dao simulation.EventDAO, initState simulation.State) {
	state := make(map[string]interface{})

	if len(initState.Stations) != 0 {
		state["counter"] = initState.Counter
		state["time"] = initState.ActualTime
		state["trains"] = initState.Trains
		state["stations"] = initState.Stations
		state["lines"] = initState.Lines
	} else {
		stations := map[string]simulation.Station{
			"MAPLE_STATION": simulation.Station{
				ID:       0,
				Location: simulation.Location{X: 1, Y: 1},
				Name:     "Maple Station",
				Riders:   []simulation.Rider{},
			},
			"PINE_STATION": simulation.Station{
				Location: simulation.Location{X: 5, Y: 1},
				Name:     "Pine Station",
				ID:       1,
				Riders:   []simulation.Rider{},
			},
			"MAHOGANY_STATION": simulation.Station{
				Location: simulation.Location{X: 8, Y: 2},
				Name:     "Mahogany Station",
				ID:       2,
				Riders:   []simulation.Rider{},
			},
			"PALM_STATION": simulation.Station{
				Location: simulation.Location{X: 14, Y: 2},
				Name:     "Palm Station",
				ID:       3,
				Riders:   []simulation.Rider{},
			},
			"ASH_STATION": simulation.Station{
				Location: simulation.Location{X: 5, Y: 4},
				Name:     "Ash Station",
				ID:       4,
				Riders:   []simulation.Rider{},
			},
			"CEDAR_STATION": simulation.Station{
				Location: simulation.Location{X: 14, Y: 5},
				Name:     "Cedar Station",
				ID:       5,
				Riders:   []simulation.Rider{},
			},
			"REDWOOD_STATION": simulation.Station{
				Location: simulation.Location{X: 2, Y: 6},
				Name:     "Redwood Station",
				ID:       6,
				Riders:   []simulation.Rider{},
			},
			"ELM_STATION": simulation.Station{
				Location: simulation.Location{X: 8, Y: 4},
				Name:     "Elm Station",
				ID:       7,
				Riders:   []simulation.Rider{},
			},
			"HOLLY_STATION": simulation.Station{
				Location: simulation.Location{X: 12, Y: 6},
				Name:     "Holly Station",
				ID:       8,
				Riders:   []simulation.Rider{},
			},
			"OAK_STATION": simulation.Station{
				Location: simulation.Location{X: 5, Y: 5},
				Name:     "Oak Station",
				ID:       9,
				Riders:   []simulation.Rider{},
			},
		}
		state["stations"] = getStationSlice(stations)
		lines := []simulation.Line{
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
				Stations: []simulation.Station{
					stations["ELM_STATION"],
					stations["HOLLY_STATION"],
					stations["CEDAR_STATION"],
				},
				Name: "Banana",
			},
		}
		state["lines"] = lines
		// for each line assign a train on it
		trains := []simulation.Train{}
		for _, line := range lines {
			trains = append(trains, simulation.Train{
				CurrentStation: line.Stations[0],
				Line:           line,
				Riders:         []simulation.Rider{},
				Capacity:       10,
				Direction:      true,
			})
		}
		state["trains"] = trains
		state["counter"] = 0
		state["time"] = time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)
	}

	reducers := []redux.Reducer{
		RiderStationReducer(dao),
		TrainStationReducer,
		RiderTrainReducer(dao),
		PersistStateReducer,
	}
	Store = redux.NewStore(state, reducers)
}

func getStationSlice(m map[string]simulation.Station) []simulation.Station {
	result := []simulation.Station{}
	for _, s := range m {
		result = append(result, s)
	}
	return result
}
