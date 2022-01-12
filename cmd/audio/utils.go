package audio

import (
	"fmt"
	"os/exec"
	"time"
)

func WaitForChannel(channel *chan bool, timeout int) bool {
	// wait 5 secs for an error to occur
	for i := 0; i < timeout; i++ {
		select {
		case <-*channel:
			return false
		default:
			fmt.Print("")
		}
		time.Sleep(time.Second)
	}
	return true
}

func killProcess(process string, channel chan bool) {
	out, err := exec.Command("killall", process).Output()
	fmt.Println(string(out))
	if err != nil {
		channel <- false
	} else {
		channel <- true
	}
}

// TODO make INIT function for logger, will be implemented in MAIN
