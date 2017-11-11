package simulation

import (
	"math/rand"
	"time"

	"github.com/mohae/deepcopy"
)

// Event represents single event in discrete event based simulation
type Event struct {
	Chance int         // chance is used to indicate the weight of certain event happening
	Value  interface{} // value is what event carries
}

// EventSimulation takes a list of events and randomly select a event to return its value
func EventSimulation(events []Event) interface{} {
	total := 0
	copyOfEvents := make([]Event, len(events))
	for i, event := range events {
		total = total + event.Chance
		copyOfEvent := deepcopy.Copy(event).(Event)
		copyOfEvent.Chance = total
		copyOfEvents[i] = copyOfEvent
	}
	randomValue := rand.Intn(total)
	for _, event := range copyOfEvents {
		if randomValue < event.Chance {
			return event.Value
		}
	}
	return nil
}

// RandomItem returns a random item from the list
func RandomItem(list []interface{}) interface{} {
	ran := rand.Intn(len(list))
	return list[ran]
}

/**
IDEA: probably move the train related struct into `metro` package and keep
simulation package purely for abstract simulation logic like EventSimulation
*/

// EventDAO is to define the serialization interface
type EventDAO interface {
	StoreRiderEvent(action, stationName, lineName string, actualTime time.Time) error
}

// State represents current application state
type State struct {
	ActualTime time.Time `json:"currentTime"`
	Counter    int64     `json:"counter"`
	Trains     []Train   `json:"trains"`
	Stations   []Station `json:"stations"`
	Lines      []Line    `json:"lines"`
}

// Location represents the geo location of the object
type Location struct {
	X int
	Y int
}

// Station contains potential riders going on at the station
type Station struct {
	Location Location
	Riders   []Rider
	Name     string
	ID       int
}

// Line represents the train line (a list of locations)
type Line struct {
	Stations []Station
	Name     string
}

// LineFilter apply filter fn to the lines
func LineFilter(lines []Line, fn func(line Line) bool) []Line {
	vsf := make([]Line, 0)
	for _, v := range lines {
		if fn(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// CastLinesToInterfaces casts the slice of lines to slice of interfaces
func CastLinesToInterfaces(lines []Line) []interface{} {
	result := make([]interface{}, len(lines))
	for i, line := range lines {
		result[i] = deepcopy.Copy(line)
	}
	return result
}

// Rider represents a single rider with its detination
type Rider struct {
	DestinationID int
}

// RiderFilter applies filter function to a list of rider
func RiderFilter(riders []Rider, fn func(rider Rider) bool) []Rider {
	vsf := make([]Rider, 0)
	for _, v := range riders {
		if fn(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Train choo choo
type Train struct {
	CurrentStation Station
	Line           Line
	Riders         []Rider
	Capacity       int
	Direction      bool // indicates either going backward or forward
}

// GetDestinations gets all destinations in order of train in this line
func (t Train) GetDestinations() []Station {
	destinations := []Station{}
	stations := t.Line.Stations
	if !t.Direction {
		stations = reverse(stations)
	}
	shouldAdd := false
	for _, station := range stations {
		if station.ID == t.CurrentStation.ID {
			shouldAdd = true
			continue
		}
		if shouldAdd {
			destinations = append(destinations, station)
		}
	}
	return destinations
}

// GetNextStation for this train on its running line
func (t Train) GetNextStation() Station {
	return t.GetDestinations()[0]
}

// Departure moves train to its next station
func (t Train) Departure() Train {
	newTrainState := deepcopy.Copy(t).(Train)
	newDirection := newTrainState.Direction
	if len(t.GetDestinations()) == 1 {
		newDirection = !newTrainState.Direction
	}
	newTrainState.CurrentStation = newTrainState.GetNextStation()
	newTrainState.Direction = newDirection

	return newTrainState
}

// CastStationsToInterfaces casts a slice of station to a slice of interfaces
func CastStationsToInterfaces(stations []Station) []interface{} {
	result := make([]interface{}, len(stations))
	for i, station := range stations {
		result[i] = deepcopy.Copy(station)
	}
	return result
}

// StationsFilter filter the station and return the stations that meets the filter Fn
func StationsFilter(stations []Station, fn func(station Station) bool) []Station {
	vsf := make([]Station, 0)
	for _, v := range stations {
		if fn(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// StationsFind finds first station meeting the fn criteria
func StationsFind(stations []Station, fn func(station Station) bool) Station {
	result := Station{}
	for _, s := range stations {
		if fn(s) {
			result = s
			break
		}
	}
	return result
}

// StationsContains return boolean to indicate if any of the station meets the criteria of fn
func StationsContains(stations []Station, fn func(station Station) bool) bool {
	for _, v := range stations {
		if fn(v) {
			return true
		}
	}
	return false
}

func reverse(stations []Station) []Station {
	result := make([]Station, len(stations))
	for i := len(stations) - 1; i >= 0; i-- {
		result[len(stations)-1-i] = stations[i]
	}
	return result
}
