package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mohae/deepcopy"
	simulation "github.com/non-player-games/metro-simulation"
	"github.com/rcliao/redux"
)

var numOfRidersGenerated = 10

// RiderStationReducer simulates the rider showing up at station
func RiderStationReducer(state redux.State, action redux.Action) redux.State {
	switch action.Type {
	case "RIDER_SHOWS_UP_STATION":
		fmt.Println("simulating riders")
		stations := deepcopy.Copy(state["stations"]).([]simulation.Station)
		// 1. generate a list of riders with their expected destination
		newRiders := []simulation.Rider{}
		for i := 0; i < numOfRidersGenerated; i++ {
			newRiders = append(
				newRiders,
				simulation.Rider{
					DestinationID: simulation.RandomItem(simulation.CastStationsToInterfaces(stations)).(simulation.Station).ID,
				},
			)
		}
		// 2. Based on destination, we will put rider into a random station in the same line
		lines := state["lines"].([]simulation.Line)
		for _, rider := range newRiders {
			linesRidersCanBe := simulation.LineFilter(lines, func(line simulation.Line) bool {
				return simulation.StationsContains(line.Stations, func(station simulation.Station) bool {
					return station.ID == rider.DestinationID
				})
			})
			if len(linesRidersCanBe) == 0 {
				log.Println("Rider doesn't belong to any line. Skipping.", rider)
				continue
			}
			lineToSendRiderTo := simulation.RandomItem(simulation.CastLinesToInterfaces(linesRidersCanBe)).(simulation.Line)
			randomStationID := rider.DestinationID
			for randomStationID == rider.DestinationID {
				randomStationID = simulation.RandomItem(simulation.CastStationsToInterfaces(lineToSendRiderTo.Stations)).(simulation.Station).ID
			}
			for i := range stations {
				if stations[i].ID == randomStationID {
					stations[i].Riders = append(stations[i].Riders, rider)
					log.Println("rider shows up at station", rider, stations[i])
				}
			}
		}
		state["stations"] = stations
		b, err := json.MarshalIndent(stations, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}
		log.Println("modified stations", string(b))
		return state
	default:
		return state
	}
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
