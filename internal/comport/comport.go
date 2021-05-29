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
	portName           string
	baudRate           int
	conn               *serial.Port
	ch                 chan []byte
	CombinedChan       chan sensordata.SensorData
	funcAssembleStatus bool //紀錄assemble function的開啟或關閉狀態
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

func (c *Comport) assemble() {
	stopchan := make(chan bool, 2)
	defer close(stopchan)

	//三秒鐘後停止這個goroutine
	go func() {
		time.Sleep(3 * time.Second)
		stopchan <- true
	}()

	str := make([]byte, 0)
	for {
		select {
		case tmp := <-c.ch:
			str = append(str, tmp...)

			//延遲合併速度，避免速度太快導致判斷為已經中斷傳輸
			time.Sleep(200 * time.Millisecond)
		case <-stopchan:
			c.funcAssembleStatus = false
			return
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
			//延遲讀取速度，避免讀取太快導致讀取到同一個數值
			time.Sleep(100 * time.Millisecond)
			n, err := c.conn.Read(buf)
			if err != nil {
				log.Println("ERROR: Failed to read data")
				c.Connect()
				continue
			}
			if c.funcAssembleStatus == false {
				//放出一個goroutine用於組裝從buffer接收到的字串
				go c.assemble()
				c.funcAssembleStatus = true
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
