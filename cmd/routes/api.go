package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MikMuellerDev/radiGo/audio"
	utils "github.com/MikMuellerDev/radiGo/utils"

	"github.com/gorilla/mux"
)

// Returns all modes / stations
func getAllModes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.GetModes())
}

// Returns the mode which is currently active
func getMode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(StatusStruct{Mode: audio.GetMode()})
}

// Returns the mode which is currently active
func getVersion(w http.ResponseWriter, r *http.Request) {
	version, instanceName, production := utils.GetVersion()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VersionStruct{Version: version, Name: instanceName, Production: production})
}

func setMode(w http.ResponseWriter, r *http.Request) {
	// If there is a start / stop operation, return an error to prevent inconsistency between shown and actual behavior
	if audio.GetOperationLock() {
		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 7, Title: "operation running", Message: "An operation is already running."})
		return
	}
	audio.SetOperationLock(true)

	success := false
	vars := mux.Vars(r)
	channel := make(chan bool)
	modePrevious := audio.GetMode()
	instruction := vars["instruction"]

	if instruction == audio.GetMode() {
		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 2, Title: "unchanged", Message: "The specified operational mode is already active."})
		audio.SetOperationLock(false)
		return
	}

	switch {
	case instruction == "off":
		audio.SetMode("off")
		audio.StopAll(4)
		success = true
	case instruction == "jellyfin":
		args := make([]string, 0)
		audio.SetMode(instruction)
		// If nothing plays, then don't attempt to kill something
		if audio.GetMode() != "off" {
			audio.StopAll(3)
			success = true
		}
		go audio.StartService("jellyfin-mpv-shim", args, channel)
		success = audio.WaitForChannel(&channel, 5)
	case utils.DoesStationExist(instruction):
		args := append(make([]string, 0), utils.GetStationById(instruction).Url, fmt.Sprintf("--volume=%d", utils.GetStationById(instruction).Volume), "--no-video")
		audio.SetMode(instruction)
		// If nothing plays, then don't attempt to kill something
		if audio.GetMode() != "off" {
			audio.StopAll(3)
		}
		go audio.StartService("mpv", args, channel)
		success = audio.WaitForChannel(&channel, 3)
	default:
		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 4, Title: "unknown mode", Message: "The specified operational mode is not valid."})
	}

	switch success {
	case true:
		json.NewEncoder(w).Encode(ResponseStruct{Success: true, ErrorCode: 0, Title: "success", Message: fmt.Sprintf("Operational mode was changed to: %s", audio.GetMode())})
	default:
		audio.SetMode(modePrevious)
		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 1, Title: "error occurred", Message: fmt.Sprintf("Something went wrong whilst trying to set mode to %s.", instruction)})
	}
	audio.SetOperationLock(false)
}
