package routes

import (
	"net/http"

	middleware "github.com/MikMuellerDev/radiGo/middleware"

	"github.com/gorilla/mux"
)

// Initializes a new Router, used in main.go
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/dash", middleware.AuthRequired(dashGetHandler)).Methods("GET")

	r.HandleFunc("/api/mode/{instruction}", middleware.AuthRequired(setMode)).Methods("POST")
	r.HandleFunc("/api/mode/list", getAllModes).Methods("GET")
	r.HandleFunc("/api/mode", getMode).Methods("GET")

	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutGetHandler).Methods("GET")

	fs := http.FileServer(http.Dir("../static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}
