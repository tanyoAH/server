package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

type Config struct {
	ApiUrl                string   `json:"apiUrl"`
	WsUrl                 string   `json:"wsUrl"`
	DBConnectionUrl       string   `json:"dbConnectionUrl"`
	DBName                string   `json:"dbName"`
	GeoLiteMMDBPath       string   `json:"geoLiteMMDBPath"`
	Home                  string   `json:"home"`
	Origin                string   `json:"origin"`
	PreferredLinkProtocol string   `json:"preferredLinkProtocol"`
	Mode                  string   `json:"mode"`
	ReservedWSENVIds      []string `json:"reservedWSENVIds"`
}

var Conf Config

const (
	DebugMode      string = "debug"
	ProductionMode string = "production"
)

func init() {
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(fmt.Sprint("Could not parse config:", err))
	}
	err = json.Unmarshal(configFile, &Conf)
	if err != nil {
		panic(fmt.Sprint("Error unmarshalling config:", err))
	}
}

func (cfg *Config) GetLogger() *logrus.Logger {
	var l = logrus.New()
	l.Formatter = &logrus.JSONFormatter{}
	return l
}

func (cfg *Config) IsDebugMode() bool {
	return cfg.Mode == DebugMode
}

func (cfg *Config) IsProductionMode() bool {
	return cfg.Mode == ProductionMode
}
