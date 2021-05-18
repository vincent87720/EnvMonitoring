package mysql

import (
	"database/sql"
	"fmt"

	"github.com/vincent87720/EnvMonitoring/internal/sensordata"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	connStr string
}

func (m *MySQL) SetConnectionString(s string) {
	m.connStr = s
}

func (m *MySQL) OpenDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", m.connStr)
	if err != nil {
		// fmt.Println(db)
		return db, err
	}

	return
}

func (m *MySQL) CreateSensorData(data sensordata.SensorData) (err error) {

	db, err := m.OpenDatabase()
	if err != nil {
		return err
	}

	_, err = db.Query("call CreateSensorData(?,?,?,?)", data.Arduino, data.Temperature, data.Humidity, data.TimeStamp)
	if err != nil {
		db.Close()
		return err
	}
	fmt.Printf("UPLOAD: %+v\n", data)
	db.Close()
	return nil
}
