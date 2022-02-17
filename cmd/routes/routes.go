package routes

import (
	"fmt"
	"net/http"

	middleware "github.com/MikMuellerDev/radiGo/middleware"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

var log *logrus.Logger

func InitLogger(logger *logrus.Logger) {
	log = logger
}

// Initializes a new Router, used in main.go
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middleware.LogRequest(indexGetHandler)).Methods("GET")
	r.HandleFunc("/dash", middleware.AuthRequired(dashGetHandler)).Methods("GET")

	r.HandleFunc("/api/mode/{instruction}", middleware.LogRequest(middleware.AuthRequired(setMode))).Methods("POST")
	r.HandleFunc("/api/mode/list", middleware.LogRequest(getAllModes)).Methods("GET")
	r.HandleFunc("/api/mode", middleware.LogRequest(getMode)).Methods("GET")
	r.HandleFunc("/api/mode/keepalive", getMode).Methods("GET")
	r.HandleFunc("/api/version", middleware.LogRequest(getVersion)).Methods("GET")

	r.HandleFunc("/login", middleware.LogRequest(loginGetHandler)).Methods("GET")
	r.HandleFunc("/login", middleware.LogRequest(loginPostHandler)).Methods("POST")
	r.HandleFunc("/logout", middleware.LogRequest(logoutGetHandler)).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	filepath := "../static"
	pathprefix := "/static/"
	fs := http.FileServer(http.Dir(filepath))
	r.PathPrefix(pathprefix).Handler(http.StripPrefix(pathprefix, fs))
	log.Debug(fmt.Sprintf("Initialized new FileServer for directory: %s. with replacement prefix: %s", filepath, pathprefix))
	log.Debug("Initialized new router.")
	return r
}
