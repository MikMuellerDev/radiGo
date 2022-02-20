package audio

import (
	"fmt"
	"os/exec"
)

var playingNow string = "off"
var operationLock bool

func StopAll(timeout int) {
	jellyfinChan := make(chan bool)
	mpvChan := make(chan bool)

	go killProcess("jellyfin-mpv-shim", jellyfinChan)
	successJellyfin := WaitForChannel(&jellyfinChan, timeout)

	go killProcess("mpv", mpvChan)
	successMpv := WaitForChannel(&mpvChan, timeout)

	if !successJellyfin || !successMpv {
		log.Trace(fmt.Sprintf("Kill jellyfin successful:%t", successJellyfin))
		log.Trace(fmt.Sprintf("Kill MPV successful:%t", successMpv))
	}

}

func StartService(command string, args []string, channel chan bool) {
	out, err := exec.Command(command, args...).Output()

	// if there is an error with our execution, handle it here
	if err != nil {
		if (GetMode() != "off" || GetMode() != "jellyfin") && !GetOperationLock() {
			log.Error(fmt.Sprintf("Error of command: %s with args: %s \n	%s", command, args, err))
		}
		channel <- false
	} else {
		channel <- true
	}
	log.Trace(fmt.Sprintf("Output of command: %s with args: %s \n	%s", command, args, string(out[:])))
}

func SetMode(mode string) {
	playingNow = mode
	log.Debug(fmt.Sprintf("Set current mode to:%s", mode))
}

func GetMode() string {
	return playingNow
}

func SetOperationLock(mode bool) {
	operationLock = mode
	log.Trace(fmt.Sprintf("OperationLock is set to:%t", mode))
}

func GetOperationLock() bool {
	return operationLock
}