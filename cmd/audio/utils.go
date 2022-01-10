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
		case out := <-*channel:
			fmt.Println("Channel data received:", out)
			return false
		default:
			fmt.Println("task: receiving data is running")
		}
		time.Sleep(time.Second)
	}
	return true
}

func killProcess(process string, channel chan bool) {
	_, err := exec.Command("killall", process).Output()
	if err != nil {
		channel <- false
	} else {
		channel <- true
	}
}
