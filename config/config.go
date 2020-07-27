package config

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	HistoricalDataSource  string `json:"history-data"`
}

var LoadedInstance Config;

func LoadConfiguration() Config {

	f, err := os.Open("./secrets/config.json")
	if err != nil {
		logrus.Error("Could not find config file or no file specified")
		logrus.Error(err)
		os.Exit(1)
	}
	defer f.Close()

	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logrus.Error(err)
	}

	LoadedInstance = cfg
	return cfg
}
