package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/non-player-games/metro-simulation/dao"
	"github.com/non-player-games/metro-simulation/store"
	"github.com/non-player-games/metro-simulation/ticker"
	"github.com/non-player-games/metro-simulation/web"
	"github.com/rcliao/redux"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db = getDB()
	mysqlDAO := dao.NewMySQLEventDAO(db)
	store.Init(mysqlDAO)
}

func main() {
	duration := 5 * time.Second
	// TODO: replace duration to 1*time.Minute
	ticker := ticker.NewTicker(duration, simulationTick(store.Store))
	ticker.Run()

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(web.Index())
	r.HandleFunc("/hello", web.Hello()).Methods("GET")
	r.HandleFunc("/api/v1/state", web.CurrentState()).Methods("GET")
	r.PathPrefix("/assets").Handler(web.Assets())

	log.Println("Running web server at port 8000")
	http.ListenAndServe(":8000", r)
}

func simulationTick(store *redux.Store) func(t time.Time) error {
	return func(t time.Time) error {
		store.Dispatch(redux.Action{Type: "TRAIN_DEPARTURE"})
		store.Dispatch(redux.Action{Type: "RIDER_SHOWS_UP_STATION"})
		store.Dispatch(redux.Action{Type: "RIDER_DEPARTURE_TRAIN"})
		store.Dispatch(redux.Action{Type: "RIDER_ARRIVAL_TRAIN"})
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
