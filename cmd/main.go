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

	// Deactivate everything
	go func() {
		audio.StopAll(5)
		audio.StartupTone()
	}()

	// Setup Scheduler to run every 6 hours
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(6).Hours().Do(audio.Reload)
	// scheduler.Every(1).Minutes().Do(audio.Reload)
	scheduler.StartAsync()

	config := utils.GetConfig()
	config.Version = "1.4.0"
	r := routes.NewRouter()
	utils.ReadModesFile()
	utils.ReadConfigFile()
	middleware.InitializeLogin(config)

	sessions.Init(config.Production)
	templates.LoadTemplates("../templates/*.html")
	http.Handle("/", r)
	log.Info(fmt.Sprintf("\x1b[34mRadiGo [Version %s] [%s] is running on http://localhost:%d", config.Version, config.InstanceName, config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
