package comport

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/vincent87720/EnvMonitoring/internal/sensordata"

	"github.com/tarm/serial"
)

type Comport struct {
	portName     string
	baudRate     int
	conn         *serial.Port
	ch           chan []byte
	CombinedChan chan sensordata.SensorData
}

func New(portName string, baudRate int) (c *Comport) {
	c = &Comport{
		portName: portName,
		baudRate: baudRate,
	}
	c.ch = make(chan []byte, 20)
	c.CombinedChan = make(chan sensordata.SensorData, 20)
	return c
}

func (c *Comport) Assemble() {

	str := make([]byte, 0)
	for {
		select {
		case tmp := <-c.ch:
			str = append(str, tmp...)
			time.Sleep(20 * time.Millisecond)
		default:
			if len(str) > 0 {
				currentTime := time.Now()
				// fmt.Println(str)

				var sd sensordata.SensorData
				json.Unmarshal(str, &sd)

				fmt.Println("GET:   ", string(str))

				sd.TimeStamp = currentTime.Format(time.RFC3339Nano)
				c.CombinedChan <- sd
				str = str[:0]
			}
		}
	}
}

func (c *Comport) Read() {

	buf := make([]byte, 128)
	for {
		for {
			n, err := c.conn.Read(buf)
			if err != nil {
				log.Println("ERROR: Failed to read data")
				c.Connect()
				continue
			}

			c.ch <- buf[:n]
			// log.Printf("%q", buf[:n])
		}
	}
}

func (c *Comport) Connect() {

	for {
		cp := &serial.Config{Name: c.portName, Baud: c.baudRate}
		s, err := serial.OpenPort(cp)
		if err != nil {
			log.Println("ERROR: Unrecognize serial port")
		} else {
			c.conn = s
			return
		}
		time.Sleep(10 * time.Second)
	}
}
