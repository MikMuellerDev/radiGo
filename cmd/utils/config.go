package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	InstanceName string
	UrlsEnabled  bool
	Production   bool
	Version      string
	Users        []User
	Port         int
}

type User struct {
	Password string
	Name     string
}

var config Config

func ReadConfigFile() {
	path := "../config/config.json"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Could not open config file: ", err)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Could not parse config file: ", err)
	}
	log.Debug(fmt.Sprintf("Loaded radiGo config File from %s", path))
	if config.UrlsEnabled {
		log.Info("Custom urls are enabled")
	}
}

func GetConfig() *Config {
	return &config
}

func GetVersion() (string, string, bool) {
	return config.Version, config.InstanceName, config.Production
}
