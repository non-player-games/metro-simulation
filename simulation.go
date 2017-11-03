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
    Name string
    ID int
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

// Define Initial Stations Need 10

var MAPLE_STATION = Station { Location: Location{1,1}, Name: "Maple Station", ID: 0 }
var PINE_STATION = Station { Location: Location{5,1}, Name: "Pine Station", ID: 1 }
var mahogany_staition = Station { Location: Location{8,2}, Name: "Mahogany Station", ID: 2 }
var Palm_station = Station { Location: Location{14,2}, Name: "Palm Station", ID: 3 }
var ash_station = Station { Location: Location{5,4}, Name: "Ash Station", ID: 4 }
var cedar_station = Station { Location: Location{14,5}, Name: "Cedar Station", ID: 5 }
var redwood_station = Station { Location: Location{2,6}, Name: "Redwood Station", ID: 6 }
var elm_station = Station { Location: Location{8,4}, Name: "Elm Station", ID: 7 }
var holly_station = Station { Location: Location{12,6}, Name: "Holly Station", ID: 8 }
var oak_station = Station { Location: Location{5,5}, Name: "Oak Station", ID: 9 }
// const STATIONS [10]Station
/*
0 Maple Station, 1,1
1 Pine Station 5,1
2 Mahogany Station 8,2
3 Palm Station 14,2
4 Ash Station 5,4
5 Cedar Station 14, 5
6 Redwood Station 2,6
7 Elm Station 8,4
8 Holly Station 12, 6
9 Oak Station 5,5
*/
