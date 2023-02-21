package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MikMuellerDev/radiGo/audio"
	"github.com/MikMuellerDev/radiGo/middleware"
	"github.com/MikMuellerDev/radiGo/routes"
	"github.com/MikMuellerDev/radiGo/sessions"
	"github.com/MikMuellerDev/radiGo/templates"
	utils "github.com/MikMuellerDev/radiGo/utils"
	"github.com/go-co-op/gocron"
)

func main() {
	log := utils.NewLogger()
	// Initialize all loggers
	audio.InitLogger(log)
	middleware.InitLogger(log)
	routes.InitLogger(log)
	sessions.InitLogger(log)
	templates.InitLogger(log)
	utils.InitLogger(log)
	log.Debug("Loggers initialized.")

	// Setup scheduler to run every 6 hours
	scheduler := gocron.NewScheduler(time.Local)
	if _, err := scheduler.Every(6).Hours().Do(audio.Reload); err != nil {
		log.Fatal("Could not start audio reload scheduler: ", err.Error())
	}
	// scheduler.Every(1).Minutes().Do(audio.Reload)
	scheduler.StartAsync()

	config := utils.GetConfig()
	config.Version = "1.5.0"
	r := routes.NewRouter()
	utils.ReadModesFile()
	utils.ReadConfigFile()
	middleware.InitializeLogin(config)
	sessions.Init(config.Production)
	templates.LoadTemplates("../templates/*.html")
	http.Handle("/", r)
	log.Info(fmt.Sprintf("\x1b[34mRadiGo [Version %s] [%s] is running on http://localhost:%d", config.Version, config.InstanceName, config.Port))

	go func() {
		var autoStartId string = ""

		for _, mode := range utils.GetModes().Modes {
			if mode.AutoStart {
				autoStartId = mode.Id
			}
		}

		audio.StopAll(5)

		if autoStartId == "" {
			log.Info("No mode is marked as auto start, not starting a stream")
			return
		}

		log.Info(fmt.Sprintf("Starting mode `%s` as it is marked as auto start...", autoStartId))

		channel := make(chan bool)

		switch {
		case autoStartId == "off":
			go audio.Beep()
			audio.SetMode("off")
			audio.StopAll(4)
		case autoStartId == "jellyfin":
			go audio.Beep()
			args := make([]string, 0)
			audio.SetMode(autoStartId)
			// If nothing plays, then don't attempt to kill something
			if audio.GetMode() != "off" {
				audio.StopAll(3)
			}
			go audio.StartService("jellyfin-mpv-shim", args, channel)
			audio.WaitForChannel(&channel, 5)
		case utils.DoesStationExist(autoStartId):
			go audio.Beep()
			args := append(make([]string, 0), utils.GetStationById(autoStartId).Url, fmt.Sprintf("--volume=%d", utils.GetStationById(autoStartId).Volume), "--no-video")
			audio.SetMode(autoStartId)
			// If nothing plays, then don't attempt to kill something
			if audio.GetMode() != "off" {
				audio.StopAll(3)
			}
			go audio.StartService("mpv", args, channel)
			audio.WaitForChannel(&channel, 3)
		}
	}()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
