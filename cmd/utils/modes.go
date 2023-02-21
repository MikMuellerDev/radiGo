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
	AutoStart   bool
}

type Modes struct {
	Modes []Mode
}

var stations Modes

func ReadModesFile() {
	path := "../config/modes.json"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Could not open modes file: ", err)
	}
	err = json.Unmarshal(content, &stations)
	if err != nil {
		log.Fatal("Could not parse modes file: ", err)
	}

	var autoStartMode *string = nil
	for _, mode := range stations.Modes {
		if autoStartMode != nil && mode.AutoStart {
			log.Fatal(fmt.Sprintf("Cannot set station `%s` as auto start: station `%s` is already marked as auto start", mode.Id, *autoStartMode))
		}
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
	log.Trace(fmt.Sprintf("Requested mode that does not exist: %s", id))
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
