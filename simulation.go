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
}

// Define Initial Stations Need 10
var MAPLE_STATION = Station{Location: Location{1, 1}, Name: "Maple Station", ID: 0}
var PINE_STATION = Station{Location: Location{5, 1}, Name: "Pine Station", ID: 1}
var MAHOGANY_STATION = Station{Location: Location{8, 2}, Name: "Mahogany Station", ID: 2}
var PALM_STATION = Station{Location: Location{14, 2}, Name: "Palm Station", ID: 3}
var ASH_STATION = Station{Location: Location{5, 4}, Name: "Ash Station", ID: 4}
var CEDAR_STATION = Station{Location: Location{14, 5}, Name: "Cedar Station", ID: 5}
var REDWOOD_STATION = Station{Location: Location{2, 6}, Name: "Redwood Station", ID: 6}
var ELM_STATION = Station{Location: Location{8, 4}, Name: "Elm Station", ID: 7}
var HOLLY_STATION = Station{Location: Location{12, 6}, Name: "Holly Station", ID: 8}
var OAK_STATION = Station{Location: Location{5, 5}, Name: "Oak Station", ID: 9}

// Define Lines
var TOMATO_LINE = Line{[]Station{MAPLE_STATION, PINE_STATION, MAHOGANY_STATION}, "Tomato"}
var AVOCADO_LINE = Line{[]Station{MAHOGANY_STATION, PALM_STATION, CEDAR_STATION}, "Avocado"}
var BLUEBERRY_LINE = Line{[]Station{MAPLE_STATION, REDWOOD_STATION, OAK_STATION, ASH_STATION, PINE_STATION}, "Tomato"}
var ORANGE_LINE = Line{[]Station{ASH_STATION, ELM_STATION, MAHOGANY_STATION}, "Orange"}
var BANANA_LINE = Line{[]Station{ELM_STATION, HOLLY_STATION, CEDAR_STATION}, "Tomato"}
