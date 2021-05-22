package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/vincent87720/EnvMonitoring/internal/sensordata"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	connStr string
	db      *sql.DB
}

func (m *MySQL) SetConnectionString(s string) {
	m.connStr = s
}

func (m *MySQL) OpenDatabase() (err error) {
	db, err := sql.Open("mysql", m.connStr)
	if err != nil {
		// fmt.Println(db)
		return err
	}
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime((30 * time.Second))

	m.db = db

	return
}

func (m *MySQL) CreateSensorData(data sensordata.SensorData) (err error) {

	conn, err := m.db.Query("call CreateSensorData(?,?,?,?)", data.Arduino, data.Temperature, data.Humidity, data.TimeStamp)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("UPLOAD: %+v\n", data)
	return nil
}
