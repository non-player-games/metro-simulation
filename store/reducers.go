package store

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/mohae/deepcopy"
	simulation "github.com/non-player-games/metro-simulation"
	"github.com/rcliao/redux"
)

var maxNumOfRidersGenerated = 100

// RiderStationReducer simulates the rider showing up at station
func RiderStationReducer(dao simulation.EventDAO) redux.Reducer {
	return func(state redux.State, action redux.Action) redux.State {
		switch action.Type {
		case "RIDER_SHOWS_UP_STATION":
			stations := deepcopy.Copy(state["stations"]).([]simulation.Station)
			// 1. generate a list of riders with their expected destination
			newRiders := []simulation.Rider{}
			numOfRidersGenerated := rand.Intn(maxNumOfRidersGenerated)
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
						if err := dao.StoreRiderEvent("ARRIVAL_STATION", stations[i].Name, lineToSendRiderTo.Name); err != nil {
							log.Println("has issue updating rider event", err)
						}
					}
				}
			}
			state["stations"] = stations
			return state
		default:
			return state
		}
	}
}

// TrainStationReducer simulates running train in the same line
func TrainStationReducer(state redux.State, action redux.Action) redux.State {
	// for each train: move it to next location in the train line
	switch action.Type {
	case "TRAIN_DEPARTURE":
		// IDEA: move train based on geo location with its accelocation rather teleporting to next station
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
func RiderTrainReducer(dao simulation.EventDAO) redux.Reducer {
	return func(state redux.State, action redux.Action) redux.State {
		// for each rider in the station with a train, decide if rider want to hop on the train
		switch action.Type {
		case "RIDER_ARRIVAL_TRAIN":
			trains := state["trains"].([]simulation.Train)
			stations := state["stations"].([]simulation.Station)
			for i, train := range trains {
				if len(train.Riders) >= train.Capacity {
					continue
				}
				for j, station := range stations {
					if train.CurrentStation.ID == station.ID {
						onboardingRiders := []simulation.Rider{}
						for _, rider := range station.Riders {
							for _, destination := range train.GetDestinations() {
								if destination.ID == rider.DestinationID {
									if len(train.Riders)+len(onboardingRiders) >= train.Capacity {
										continue
									}
									onboardingRiders = append(onboardingRiders, rider)
									if err := dao.StoreRiderEvent("ARRIVAL_TRAIN", station.Name, train.Line.Name); err != nil {
										log.Println("failed to insert arrival train event", err)
									}
								}
							}
						}
						for _, onboardingRider := range onboardingRiders {
							trains[i].Riders = append(trains[i].Riders, onboardingRider)
							stations[j].Riders = simulation.RiderFilter(
								stations[j].Riders,
								func(rider simulation.Rider) bool {
									for _, destination := range train.GetDestinations() {
										if rider.DestinationID == destination.ID {
											return false
										}
									}
									return true
								},
							)
						}
					}
				}
			}
			return state
		case "RIDER_DEPARTURE_TRAIN":
			trains := state["trains"].([]simulation.Train)
			for i := range trains {
				for _ = range trains[i].Riders {
					if err := dao.StoreRiderEvent("DEPARTURE_TRAIN", trains[i].CurrentStation.Name, trains[i].Line.Name); err != nil {
						log.Println("failed to insert departure train event", err)
					}
				}
				trains[i].Riders = simulation.RiderFilter(
					trains[i].Riders,
					func(rider simulation.Rider) bool {
						if trains[i].CurrentStation.ID == rider.DestinationID {
							return false
						}
						return true
					},
				)
			}
			return state
		default:
			return state
		}
	}
}

// PersistStateReducer take the state and store into JSON
func PersistStateReducer(state redux.State, action redux.Action) redux.State {
	switch action.Type {
	case "PERSIST_STATE":
		currentState := simulation.State{
			Trains:   state["trains"].([]simulation.Train),
			Stations: state["stations"].([]simulation.Station),
			Lines:    state["lines"].([]simulation.Line),
		}
		stateJSON, _ := json.Marshal(currentState)
		err := ioutil.WriteFile("state.json", stateJSON, 0644)
		if err != nil {
			log.Println("failed to write state to json")
		}
		return state
	default:
		return state
	}
}
