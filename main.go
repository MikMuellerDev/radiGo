package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var USERNAME string = "mik"
var PASSWORD string = "mik"

var store = sessions.NewCookieStore([]byte("Replace me0 12"))
var templates *template.Template

func main() {
	// Specify the template's directory
	templates = template.Must(template.ParseGlob("templates/*.html"))
	// Initialize a new Mux Router
	r := mux.NewRouter()

	// Dashboard !TODO shows all music
	r.HandleFunc("/", indexGetHandler).Methods("GET")

	r.HandleFunc("/api/jellyfin", startJellyfin).Methods("GET")
	r.HandleFunc("/api/off", stopAll).Methods("GET")

	// r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	r.HandleFunc("/dash", dashGetHandler).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	fmt.Println("♬ RadiGo is running on  http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
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

	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		templates.ExecuteTemplate(w, "login.html", nil)
		return
	}
	username, ok := untyped.(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if username == "mik" {
		http.Redirect(w, r, "/dash", http.StatusFound)
	} else {
		templates.ExecuteTemplate(w, "login.html", nil)
	}
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if username == USERNAME && password == PASSWORD {
		session, _ := store.Get(r, "session")
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/dash", http.StatusFound)
	} else {
		templates.ExecuteTemplate(w, "login.html", http.StatusForbidden)
	}
}

func dashGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	username, ok := untyped.(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// A user is logged in
	templates.ExecuteTemplate(w, "dash.html", nil)
	fmt.Println("Username:", username)
	// w.Write([]byte(username))
}

func startJellyfin(w http.ResponseWriter, r *http.Request) {
	channel := make(chan bool)
	go startJellyfinCommand(channel)

	// If the Goroutine (Jellyfin) does not complete (crash) in 10 seconds after its initialization, it is considered running
	for i := 0; i < 10; i++ {
		select {
		case <-channel:
			fmt.Println("Error starting Jellyfin")
			w.Write([]byte("Error starting Jellyfin"))
			return
		default:
			fmt.Println("Jellyfin is running")
		}
		time.Sleep(time.Second)
	}
	w.Write([]byte("Jellyfin is running"))
}

func startJellyfinCommand(channel chan bool) {
	_, err := exec.Command("jellyfin-mpv-shim").Output()

	// if there is an error with our execution, handle it here
	if err != nil {
		fmt.Printf("%s", err)
		channel <- false
	} else {
		channel <- true
	}
	// output := string(out[:])
	// fmt.Println(output)
}

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

func killJellyfinCommand(channel chan bool) {
	_, err := exec.Command("killall", "jellyfin-mpv-shim").Output()

	// if there is an error with our execution, handle it here
	if err != nil {
		channel <- false
	} else {
		channel <- true
	}
}

func waitForChannel(channel *chan bool) bool {
	// wait 5 secs for channel
	for i := 0; i < 5; i++ {
		select {
		case out := <-*channel:
			fmt.Println("Channel data received:", out)
			if out {
				return true
			}
			return false
		default:
			fmt.Println("task: receiving data is running")
		}
		time.Sleep(time.Second)
	}
	return false
}

func stopAll(w http.ResponseWriter, r *http.Request) {
	jellyfinChan := make(chan bool)

	go killJellyfinCommand(jellyfinChan)
	success := waitForChannel(&jellyfinChan)

	fmt.Println(success, "is success of killJellyfin")

	if success {
		w.Write([]byte("Success"))
	} else {
		w.Write([]byte("Failure"))
	}
}
