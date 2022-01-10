package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Station struct {
	Name        string
	Description string
	ImagePath   string
	Url         string
	Id          string
}

type Stations struct {
	Stations []Station
}

var stations Stations

func InitStations() {
	content, err := ioutil.ReadFile("../config/stations.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Stations
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	fmt.Println(payload)
	stations = payload
}

func GetStations() Stations {
	return stations
}

func DoesStationExist(id string) bool {
	for _, v := range stations.Stations {
		if v.Id == id {
			return true
		}
	}
	return false
}

func GetStationById(id string) Station {
	for _, v := range stations.Stations {
		if v.Id == id {
			return v
		}
	}
	return Station{}
}
