package main

import (
	"fmt"
	"net/http"

	"github.com/MikMuellerDev/radiGo/audio"
	"github.com/MikMuellerDev/radiGo/middleware"
	"github.com/MikMuellerDev/radiGo/routes"
	"github.com/MikMuellerDev/radiGo/sessions"
	"github.com/MikMuellerDev/radiGo/templates"
	utils "github.com/MikMuellerDev/radiGo/utils"
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
	log.Debug("All loggers initialized.")

	config := utils.GetConfig()
	r := routes.NewRouter()
	utils.ReadModesFile()
	utils.ReadConfigFile()
	middleware.InitializeLogin(config)

	sessions.Init(config.Production)
	templates.LoadTemplates("../templates/*.html")
	http.Handle("/", r)
	log.Info(fmt.Sprintf("\x1b[34mRadiGo [%s] is running on http://localhost:%d", config.InstanceName, config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
