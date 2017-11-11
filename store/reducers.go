package store

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	simulation "github.com/non-player-games/metro-simulation"
	"github.com/rcliao/redux"
)

var initialActualTime = time.Date(2017, time.January, 1, 0, 0, 0, 0, time.Local)
var actualDuration = time.Hour
var hourToRidersMapping = map[int]int{
	0:  20,
	1:  30,
	2:  20,
	3:  20,
	4:  30,
	5:  40,
	6:  90,
	7:  130,
	8:  200,
	9:  700,
	10: 300,
	11: 350,
	12: 500,
	13: 550,
	14: 330,
	15: 430,
	16: 530,
	17: 630,
	18: 430,
	19: 330,
	20: 150,
	21: 50,
	22: 40,
	23: 30,
}
var stationIDsPercentage = []simulation.Event{
	simulation.Event{
		Chance: 5,
		Value:  0,
	},
	simulation.Event{
		Chance: 30,
		Value:  1,
	}, simulation.Event{
		Chance: 30,
		Value:  2,
	}, simulation.Event{
		Chance: 10,
		Value:  3,
	}, simulation.Event{
		Chance: 20,
		Value:  4,
	}, simulation.Event{
		Chance: 5,
		Value:  5,
	}, simulation.Event{
		Chance: 3,
		Value:  6,
	}, simulation.Event{
		Chance: 50,
		Value:  7,
	}, simulation.Event{
		Chance: 3,
		Value:  8,
	}, simulation.Event{
		Chance: 20,
		Value:  9,
	},
}

// RiderStationReducer simulates the rider showing up at station
func RiderStationReducer(dao simulation.EventDAO) redux.Reducer {
	return func(state redux.State, action redux.Action) redux.State {
		switch action.Type {
		case "RIDER_SHOWS_UP_STATION":
			stations := state["stations"].([]simulation.Station)
			logicalTime := action.Value.(int64)
			simulatedTime := logicalTimeToRealTime(initialActualTime, actualDuration, logicalTime)
			// 1. generate a list of riders on a station
			numOfRidersGenerated := rand.Intn(getMaximumRiderBasedOnHour(simulatedTime.Hour()))
			numOfRidersPerStation := make(map[int]int, numOfRidersGenerated)
			for i := 0; i < numOfRidersGenerated; i++ {
				stationID := simulation.EventSimulation(stationIDsPercentage).(int)
				numOfRidersPerStation[stationID]++
			}
			// 2. Based on station, we will give this rider into a random destination in the same line
			lines := state["lines"].([]simulation.Line)
			for stationID, numberOfRiders := range numOfRidersPerStation {
				linesRidersCanBe := simulation.LineFilter(lines, func(line simulation.Line) bool {
					return simulation.StationsContains(line.Stations, func(station simulation.Station) bool {
						return station.ID == stationID
					})
				})
				if len(linesRidersCanBe) == 0 {
					log.Println("Rider doesn't belong to any line. Skipping.", stationID)
					continue
				}
				for r := 0; r < numberOfRiders; r++ {
					lineToSendRiderTo := simulation.RandomItem(simulation.CastLinesToInterfaces(linesRidersCanBe)).(simulation.Line)
					randomStationID := stationID
					for randomStationID == stationID {
						randomStationID = simulation.RandomItem(simulation.CastStationsToInterfaces(lineToSendRiderTo.Stations)).(simulation.Station).ID
					}
					rider := simulation.Rider{DestinationID: randomStationID}
					for i := range stations {
						if stations[i].ID == stationID {
							stations[i].Riders = append(stations[i].Riders, rider)
							log.Println("rider shows up at station", rider, stations[i])
							if err := dao.StoreRiderEvent("ARRIVAL_STATION", stations[i].Name, lineToSendRiderTo.Name, simulatedTime); err != nil {
								log.Println("has issue updating rider event", err)
							}
							break
						}
					}
				}
			}
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
			logicalTime := action.Value.(int64)
			simulatedTime := logicalTimeToRealTime(initialActualTime, actualDuration, logicalTime)
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
										if err := dao.StoreRiderEvent("TRAIN_FULL", station.Name, train.Line.Name, simulatedTime); err != nil {
											log.Println("failed to insert train full event", err)
										}
										continue
									}
									onboardingRiders = append(onboardingRiders, rider)
									if err := dao.StoreRiderEvent("ARRIVAL_TRAIN", station.Name, train.Line.Name, simulatedTime); err != nil {
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
			logicalTime := action.Value.(int64)
			simulatedTime := logicalTimeToRealTime(initialActualTime, actualDuration, logicalTime)
			trains := state["trains"].([]simulation.Train)
			for i := range trains {
				for _ = range trains[i].Riders {
					if err := dao.StoreRiderEvent("DEPARTURE_TRAIN", trains[i].CurrentStation.Name, trains[i].Line.Name, simulatedTime); err != nil {
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
		logicalTime := action.Value.(int64)
		actualTime := logicalTimeToRealTime(initialActualTime, actualDuration, logicalTime)
		state["counter"] = logicalTime
		state["time"] = actualTime
		currentState := simulation.State{
			Counter:    logicalTime,
			ActualTime: actualTime,
			Trains:     state["trains"].([]simulation.Train),
			Stations:   state["stations"].([]simulation.Station),
			Lines:      state["lines"].([]simulation.Line),
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

// private helper method to convert logical time based on initial time and duration
// IDEA: move this to ticker package
func logicalTimeToRealTime(initialTime time.Time, duration time.Duration, logicalTime int64) time.Time {
	return initialTime.Add(time.Duration(logicalTime) * duration)
}

func getMaximumRiderBasedOnHour(hour int) int {
	// IDEA: would be better to replace this function as a math formula.
	// I cant think of good math formula to distribute easily so I create a hard-coded map for now
	return hourToRidersMapping[hour]
}
