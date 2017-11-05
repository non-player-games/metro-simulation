package dao

import "database/sql"

// MySQLEventDAO implements EventDAO interface from simulation to store events
type MySQLEventDAO struct {
	db *sql.DB
}

// NewMySQLEventDAO is constructor for creating new mysql event dao
func NewMySQLEventDAO(db *sql.DB) *MySQLEventDAO {
	return &MySQLEventDAO{db}
}

// StoreRiderEvent stores the rider event into mysql table
func (m *MySQLEventDAO) StoreRiderEvent(action, stationName, lineName string) error {
	stmt, err := m.db.Prepare("INSERT INTO rider_events (action, station, line) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(action, stationName, lineName)
	if err != nil {
		return err
	}
	return nil
}
