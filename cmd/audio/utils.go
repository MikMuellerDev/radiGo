package audio

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(logger *logrus.Logger) {
	log = logger
}

func WaitForChannel(channel *chan bool, timeout int) bool {
	// wait 5 secs for a potential error to occur
	for i := 0; i < timeout; i++ {
		select {
		case <-*channel:
			log.Trace("Received signal from channel")
			return false
		default:
			log.Trace("Waiting for channel")
		}
		time.Sleep(time.Second)
	}
	return true
}

func killProcess(process string, channel chan bool) {
	out, err := exec.Command("killall", process).Output()
	log.Debug(fmt.Sprintf("Output: %s", out))
	if err != nil {
		channel <- false
	} else {
		channel <- true
	}
}

// TODO make INIT function for logger, will be implemented in MAIN
