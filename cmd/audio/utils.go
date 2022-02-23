package audio

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/MikMuellerDev/radiGo/utils"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(logger *logrus.Logger) {
	log = logger
}

func WaitForChannel(channel *chan bool, timeout int) bool {
	// wait {timeout} secs for a potential error to occur
	// Returns true for success
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

func Reload() {
	if GetMode() != "off" {
		log.Debug("Reloading current mode...")
		previousMode := playingNow
		for GetOperationLock() {
			time.Sleep(1 * time.Second)
		}
		SetOperationLock(true)
		SetMode("off")
		StopAll(5)
		channel := make(chan bool)
		if GetMode() == "jellyfin" {
			args := make([]string, 0)
			go StartService("jellyfin-mpv-shim", args, channel)
			success := WaitForChannel(&channel, 5)
			if success {
				log.Info("Reloading current mode succeeded.")
			} else {
				log.Error("Reloading current mode failed.")
			}
		} else {
			args := append(make([]string, 0), utils.GetStationById(previousMode).Url, fmt.Sprintf("--volume=%d", utils.GetStationById(previousMode).Volume), "--no-video")
			go StartService("mpv", args, channel)
			success := WaitForChannel(&channel, 3)
			if success {
				log.Info("Reloading current mode succeeded.")
			} else {
				log.Error("Reloading current mode failed.")
			}
		}
		SetMode(previousMode)
		SetOperationLock(false)
	}
}
