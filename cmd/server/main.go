package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	simulation "github.com/non-player-games/metro-simulation"
	"github.com/non-player-games/metro-simulation/dao"
	"github.com/non-player-games/metro-simulation/store"
	"github.com/non-player-games/metro-simulation/ticker"
	"github.com/non-player-games/metro-simulation/web"
	"github.com/rcliao/redux"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// IDEA: probably better ro move logical time toward Ticket package
var logicalTime int64

func init() {
	db = getDB()
	mysqlDAO := dao.NewMySQLEventDAO(db)

	// read state.json if it exist and use it as initial state
	file, e := ioutil.ReadFile("state.json")
	if e != nil {
		fmt.Printf("cannot read json file. Assume state doesn't exist. %v\n", e)
	}
	var initState simulation.State
	json.Unmarshal(file, &initState)
	logicalTime = initState.Counter
	store.Init(mysqlDAO, initState)
}

func main() {
	duration := 5 * time.Second
	if os.Getenv("ENVIRONMENT") == "production" {
		duration = 1 * time.Minute
	}
	ticker := ticker.NewTicker(duration, simulationTick(store.Store))
	ticker.Run()

	r := mux.NewRouter()
	r.PathPrefix("/assets").Handler(web.Assets())
	r.HandleFunc("/hello", web.Hello()).Methods("GET")
	r.HandleFunc("/api/v1/state", web.CurrentState()).Methods("GET")
	r.PathPrefix("/").Handler(web.Index())

	log.Println("Running web server at port 8000")
	http.ListenAndServe(":8000", r)
}

func simulationTick(store *redux.Store) func(t time.Time) error {
	return func(t time.Time) error {
		logicalTime++
		store.Dispatch(redux.Action{Type: "TRAIN_DEPARTURE", Value: logicalTime})
		store.Dispatch(redux.Action{Type: "RIDER_SHOWS_UP_STATION", Value: logicalTime})
		store.Dispatch(redux.Action{Type: "RIDER_DEPARTURE_TRAIN", Value: logicalTime})
		store.Dispatch(redux.Action{Type: "RIDER_ARRIVAL_TRAIN", Value: logicalTime})
		store.Dispatch(redux.Action{Type: "PERSIST_STATE", Value: logicalTime})
		return nil
	}
}

func getDB() *sql.DB {
	defaultProtocol := "tcp"
	defaultPort := "3306"
	username := os.Getenv("MYSQL_USERNAME")
	if username == "" {
		username = "root"
	}
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	if database == "" {
		database = "metro"
	}

	sqlDSN := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s",
		username,
		password,
		defaultProtocol,
		host,
		defaultPort,
		database,
	)

	db, err := sql.Open("mysql", sqlDSN)
	if err != nil {
		panic(err)
	}

	return db
}
