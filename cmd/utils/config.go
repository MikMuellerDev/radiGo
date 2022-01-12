package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port         int
	Users        []User
	InstanceName string
}

type User struct {
	Name     string
	Password string
}

var config Config

func ReadConfigFile() {
	content, err := ioutil.ReadFile("../config/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}

func GetConfig() *Config {
	return &config
}
