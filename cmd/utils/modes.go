package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	path := "../config/modes.json"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	err = json.Unmarshal(content, &stations)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	log.Debug(fmt.Sprintf("Loaded radiGo modes and stations from %s", path))
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
	log.Trace(fmt.Sprintf("Requested mode that does not exist %s", id))
	return false
}

func GetStationById(id string) Mode {
	for _, v := range stations.Modes {
		if v.Id == id {
			return v
		}
	}
	log.Error(fmt.Sprintf("Requested station that does not exist: %s, DoesStationExist() might have failed.", id))
	return Mode{}
}
