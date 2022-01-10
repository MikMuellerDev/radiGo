package main

import (
	"fmt"
	"net/http"

	"github.com/MikMuellerDev/radiGo/routes"
	"github.com/MikMuellerDev/radiGo/templates"
	utils "github.com/MikMuellerDev/radiGo/utils"
)

func main() {
	log := utils.NewLogger()
	r := routes.NewRouter()
	utils.InitStations()
	templates.LoadTemplates("../templates/*.html")
	http.Handle("/", r)
	log.Info("♬ RadiGo is running on  http://localhost:8080")
	fmt.Println(http.ListenAndServe(":1234", nil))
}
