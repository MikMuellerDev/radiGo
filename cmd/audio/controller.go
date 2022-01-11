package audio

import (
	"fmt"
	"os/exec"
)

var playingNow string = "off"

func StopAll(timeout int) bool {
	jellyfinChan := make(chan bool)
	mpvChan := make(chan bool)

	go killProcess("jellyfin-mpv-shim", jellyfinChan)
	successJellyfin := WaitForChannel(&jellyfinChan, timeout)

	go killProcess("mpv", mpvChan)
	successMpv := WaitForChannel(&mpvChan, timeout)

	fmt.Println(successJellyfin, "is success of killJellyfin")
	fmt.Println(successMpv, "is success of Kill MPV")

	return successJellyfin || successMpv
}

func StartService(command string, arg string, channel chan bool) {
	_, err := exec.Command(command, arg).Output()

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

func SetPlaying(mode string) {
	playingNow = mode
}

func GetPlaying() string {
	return playingNow
}
