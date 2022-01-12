package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Mode struct {
	Name        string
	Description string
	ImagePath   string
	Url         string
	Id          string
	Volume      int
}

type Modes struct {
	Modes []Mode
}

var stations Modes

func ReadModesFile() {
	content, err := ioutil.ReadFile("../config/modes.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	err = json.Unmarshal(content, &stations)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}

func GetModes() Modes {
	return stations
}

func DoesStationExist(id string) bool {
	for _, v := range stations.Modes {
		if v.Id == id {
			return true
		}
	}
	return false
}

func GetStationById(id string) Mode {
	for _, v := range stations.Modes {
		if v.Id == id {
			return v
		}
	}
	return Mode{}
}
