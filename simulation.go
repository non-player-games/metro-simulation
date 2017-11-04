package simulation

import "github.com/mohae/deepcopy"

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

// Rider represents a single rider with its detination
type Rider struct {
	Destination Station
}

// Train choo choo
type Train struct {
	CurrentStation Station
	Line           Line
	Riders         []Rider
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
	newDirection := false
	if len(t.GetDestinations()) == 1 {
		newDirection = !newTrainState.Direction
	}
	newTrainState.CurrentStation = newTrainState.GetNextStation()
	newTrainState.Direction = newDirection

	return newTrainState
}

func reverse(stations []Station) []Station {
	result := make([]Station, len(stations))
	for i := len(stations) - 1; i >= 0; i-- {
		result[len(stations)-1-i] = stations[i]
	}
	return result
}
