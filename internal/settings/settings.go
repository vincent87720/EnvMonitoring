package settings

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type T struct {
	Port struct {
		Name     string `yaml:"name"`
		BaudRate int    `yaml:"baudRate"`
		DataBits int    `yaml:"dataBits"`
		Parity   string `yaml:"parity"`
		StopBits int    `yaml:"stopBits"`
	}
	Database struct {
		Host     string `yaml:"host"`
		DBname   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

type Settings struct {
	yamlByteXi []byte
	t          T
}

func (s *Settings) ReadFile() {
	ya, err := ioutil.ReadFile("./settings.yaml")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
		return
	}
	s.yamlByteXi = ya
}

func (s *Settings) UnmarshalSettings() {

	s.t = T{}
	err := yaml.Unmarshal(s.yamlByteXi, &s.t)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	// fmt.Printf("settings: %+v\n", s.t)
}

func (s *Settings) GetPortName() string {
	return s.t.Port.Name
}

func (s *Settings) GetBaudRate() int {
	return s.t.Port.BaudRate
}

func (s *Settings) GetDBConnectionString() string {
	return s.t.Database.Username + ":" + s.t.Database.Password + "@tcp(" + s.t.Database.Host + ")/" + s.t.Database.DBname + "?charset=utf8"
}
