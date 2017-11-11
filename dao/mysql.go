package dao

import (
	"database/sql"
	"time"
)

// MySQLEventDAO implements EventDAO interface from simulation to store events
type MySQLEventDAO struct {
	db *sql.DB
}

// NewMySQLEventDAO is constructor for creating new mysql event dao
func NewMySQLEventDAO(db *sql.DB) *MySQLEventDAO {
	return &MySQLEventDAO{db}
}

// StoreRiderEvent stores the rider event into mysql table
func (m *MySQLEventDAO) StoreRiderEvent(action, stationName, lineName string, actualtime time.Time) error {
	stmt, err := m.db.Prepare("INSERT INTO rider_events (action, station, line, timestamp) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(action, stationName, lineName, actualtime)
	if err != nil {
		return err
	}
	return nil
}
