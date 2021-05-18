package db

import (
	"github.com/vincent87720/EnvMonitoring/internal/sensordata"
)

type Database interface {
	CreateSensorData(data sensordata.SensorData) (err error)
}

func CreateSensorData(db Database, data sensordata.SensorData) (err error) {
	return db.CreateSensorData(data)
}
