package simulation

// Location represents the geo location of the object
type Location struct {
	X int
	Y int
}

// Station contains potential riders going on at the station
type Station struct {
	Location Location
	Riders   []Rider
}

// Line represents the train line (a list of locations)
type Line struct {
	Stations []Station
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
}
