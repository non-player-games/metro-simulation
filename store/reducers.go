package store

import (
	simulation "github.com/non-player-games/metro-simulation"
	"github.com/rcliao/redux"
)

// RiderStationReducer simulates the rider showing up at station
func RiderStationReducer(state redux.State, action redux.Action) redux.State {
	// 1. generate a list of riders with their expected destination
	// 2. Based on destination, we will put rider into a station in the same line
	return state
}

// TrainStationReducer simulates running train in the same line
func TrainStationReducer(state redux.State, action redux.Action) redux.State {
	// for each train: move it to next location in the train line
	switch action.Type {
	case "TRAIN_DEPARTURE":
		trains := state["trains"].([]simulation.Train)
		for i, train := range trains {
			trains[i] = train.Departure()
		}
		return state
	default:
		return state
	}
}

// RiderTrainReducer simulates rider deciding to hop on the train or not based on its destination
func RiderTrainReducer(state redux.State, action redux.Action) redux.State {
	// for each rider in the station with a train, decide if rider want to hop on the train
	return state
}
