package main

import (
	"fmt"
	"net/http"

	"github.com/MikMuellerDev/radiGo/middleware"
	"github.com/MikMuellerDev/radiGo/routes"
	"github.com/MikMuellerDev/radiGo/templates"
	utils "github.com/MikMuellerDev/radiGo/utils"
)

func main() {
	log := utils.NewLogger()
	config := utils.GetConfig()
	r := routes.NewRouter()
	utils.ReadModesFile()
	utils.ReadConfigFile()
	middleware.InitializeLogin(config)
	templates.LoadTemplates("../templates/*.html")
	http.Handle("/", r)
	log.Info(fmt.Sprintf("RadiGo [%s] is running on http://localhost:%d", config.InstanceName, config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
