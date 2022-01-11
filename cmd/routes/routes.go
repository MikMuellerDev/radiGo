package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MikMuellerDev/radiGo/audio"
	middleware "github.com/MikMuellerDev/radiGo/middleware"
	sessions "github.com/MikMuellerDev/radiGo/sessions"
	templates "github.com/MikMuellerDev/radiGo/templates"
	utils "github.com/MikMuellerDev/radiGo/utils"

	"github.com/gorilla/mux"
)

var USERNAME string = "mik"
var PASSWORD string = "test"

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Dashboard !TODO shows all music
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/dash", middleware.AuthRequired(dashGetHandler)).Methods("GET")

	r.HandleFunc("/stations", getStations).Methods("GET")
	r.HandleFunc("/api/mode/{instruction}", setPlaying).Methods("POST", "GET")
	// r.HandleFunc("/api/off", middleware.AuthRequired(stopAll)).Methods("POST", "GET")

	// r.HandleFunc("/api/jellyfin", middleware.AuthRequired(startJellyfin)).Methods("GET")

	// r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutGetHandler).Methods("GET")

	// r.HandleFunc("/dash", dashGetHandler).Methods("GET")

	fs := http.FileServer(http.Dir("../static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	return r
}

// Routes

func getStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(utils.GetStations())
}

// func stopAll(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	success := audio.StopAll(5)
// 	if success {
// 		json.NewEncoder(w).Encode(ResponseStruct{Success: true, Title: "success", Message: "All (music) processes were killed successfully."})
// 	} else {
// 		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 1, Title: "failure", Message: "An error occurred during the attempt to stop all (music) processes."})
// 	}
// }

func setPlaying(w http.ResponseWriter, r *http.Request) {
	// If the Goroutine (Jellyfin) does not complete (crash) in 5 seconds after its initialization, it is considered running
	vars := mux.Vars(r)
	instruction := vars["instruction"]
	modePrevious := audio.GetPlaying()
	channel := make(chan bool)
	var success bool
	fmt.Println("order: ", instruction, "current: ", audio.GetPlaying())
	if instruction == audio.GetPlaying() {
		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 2, Title: "unchanged", Message: "The specified operational mode is already active."})
		return
	}
	if instruction == "off" {
		audio.SetPlaying("off")
		success = audio.StopAll(5)
	}
	if instruction == "jellyfin" {
		audio.SetPlaying(instruction)
		// If nothing plays, then don't attempt to kill something
		if audio.GetPlaying() != "off" {
			audio.StopAll(3)
		}
		go audio.StartService("jellyfin-mpv-shim", "", channel)
		success = audio.WaitForChannel(&channel, 5)
	} else if utils.DoesStationExist(instruction) {
		audio.SetPlaying(instruction)
		// If nothing plays, then don't attempt to kill something
		if audio.GetPlaying() != "off" {
			audio.StopAll(3)
		}
		go audio.StartService("mpv", utils.GetStationById(instruction).Url, channel)
		success = audio.WaitForChannel(&channel, 5)
	} else {
		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 4, Title: "unknown mode", Message: "The specified operational mode is not valid."})
	}
	if success {
		json.NewEncoder(w).Encode(ResponseStruct{Success: true, ErrorCode: 0, Title: "success", Message: fmt.Sprintf("Operational mode was changed to: %s", audio.GetPlaying())})
	} else {
		if instruction != "off" {
			audio.SetPlaying(modePrevious)
		}
		json.NewEncoder(w).Encode(ResponseStruct{Success: false, ErrorCode: 1, Title: "error occurred", Message: fmt.Sprintf("Something went wrong whilst trying to set mode to %s.", instruction)})
	}
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dash", http.StatusFound)
}

// func indexPostHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	comment := r.PostForm.Get("comment")
// 	fmt.Println(comment)
// 	http.Redirect(w, r, "/", 302)
// }

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	value, ok := session.Values["valid"]
	valid, okParse := value.(bool)

	if ok && okParse && valid {
		http.Redirect(w, r, "/dash", http.StatusFound)
		return
	}

	fmt.Println("Value", value)
	templates.ExecuteTemplate(w, "login.html", http.StatusOK)

	// templates.ExecuteTemplate(w, "login.html", http.StatusOK)
	// if !ok {
	// 	return
	// }
	// if !valid {
	// 	templates.ExecuteTemplate(w, "login.html", http.StatusForbidden)
	// 	return
	// }
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	session.Values["valid"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if username == USERNAME && password == PASSWORD {
		session, _ := sessions.Store.Get(r, "session")
		session.Values["valid"] = true
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/dash", http.StatusFound)
	} else {
		templates.ExecuteTemplate(w, "login.html", http.StatusForbidden)
	}
}

func dashGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dash.html", http.StatusOK)
}

// func startJellyfin(w http.ResponseWriter, r *http.Request) {
// 	channel := make(chan bool)
// 	go audio.StartService("jellyfin-mpv-shim", "", channel)
// 	// go audio.StartService("mpv", "https://sunsl.streamabc.net/sunsl-melodictechno-mp3-192-4167248?sABC=61o5rrnp%230%232n97q280268772p52p03osn832218115%23ubzrcntr&mode=preroll&aw_0_1st.skey=1639311208&cb=595675563&amsparams=playerid:homepage;skey:1639313068", channel)

// 	// If the Goroutine (Jellyfin) does not complete (crash) in 10 seconds after its initialization, it is considered running
// 	for i := 0; i < 10; i++ {
// 		select {
// 		case <-channel:
// 			fmt.Println("Error starting Jellyfin")
// 			w.Write([]byte("Error starting Jellyfin"))
// 			return
// 		default:
// 			fmt.Println("Jellyfin is running")
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	w.Write([]byte("Jellyfin is running"))
// }

// func killJellyfin(w http.ResponseWriter, r *http.Request) {
// 	channel := make(chan bool)
// 	go killJellyfinCommand(channel)

// 	// If the Goroutine (Jellyfin) does not complete (crash) in 10 seconds after its initialization, it is considered running

// 	for i := 0; i < 10; i++ {
// 		select {
// 		case <-channel:
// 			fmt.Println("Jellyfin is stopped")
// 			w.Write([]byte("Jellyfin is stopped"))
// 			return
// 		default:
// 			fmt.Println("task: stopping Jellyfin is running")
// 		}
// 		time.Sleep(time.Second)
// 	}
// 	w.Write([]byte("Error stopping jellyfin, it is still running"))
// }
