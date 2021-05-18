package main

import (
	"fmt"
	"log"

	db "github.com/vincent87720/EnvMonitoring/internal/I_db"
	receiver "github.com/vincent87720/EnvMonitoring/internal/I_receiver"
	"github.com/vincent87720/EnvMonitoring/internal/comport"
	"github.com/vincent87720/EnvMonitoring/internal/mysql"
	"github.com/vincent87720/EnvMonitoring/internal/settings"
)

func main() {

	fmt.Println("Loading settings...")
	//Read settings from the settings.yaml file
	s := &settings.Settings{}
	s.ReadFile()
	s.UnmarshalSettings()
	dbconstr := s.GetDBConnectionString()

	fmt.Println("Opening serial port...")
	//以portName和baudRate建立一個Comport
	com := comport.New(s.GetPortName(), s.GetBaudRate())

	//放出一個goroutine用於組裝從buffer接收到的字串
	go com.Assemble()

	//與Serial port進行連線
	receiver.Connect(com)

	//放出一個goroutine用於監聽serial port，負責接收資料
	go receiver.Read(com)

	//測試資料庫連線
	mydb := &mysql.MySQL{}
	mydb.SetConnectionString(dbconstr)
	_, err := mydb.OpenDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Ready")

	for {
		select {
		case sd := <-com.CombinedChan:
			err := db.CreateSensorData(mydb, sd)
			if err != nil {
				log.Println(err.Error())
			}
			fmt.Println("---")
		}
	}
}
