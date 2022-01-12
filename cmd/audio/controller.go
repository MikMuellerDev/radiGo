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
		fmt.Println(successJellyfin, "is success of killJellyfin")
		fmt.Println(successMpv, "is success of Kill MPV")
	}

}

func StartService(command string, args []string, channel chan bool) {
	_, err := exec.Command(command, args...).Output()

	// if there is an error with our execution, handle it here
	if err != nil {
		fmt.Printf("%s", err)
		channel <- false
	} else {
		channel <- true
	}
	// output := string(out[:])
	// fmt.Println(output)
}

func SetMode(mode string) {
	playingNow = mode
}

func GetMode() string {
	return playingNow
}

func SetOperationLock(mode bool) {
	operationLock = mode
}

func GetOperationLock() bool {
	return operationLock
}
