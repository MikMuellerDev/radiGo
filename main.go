package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var USERNAME string = "mik"
var PASSWORD string = "mik"

var store = sessions.NewCookieStore([]byte("Replace me 2"))
var templates *template.Template

func main() {
	// Specify the template's directory
	templates = template.Must(template.ParseGlob("templates/*.html"))
	// Initialize a new Mux Router
	r := mux.NewRouter()

	// Dashboard !TODO shows all music
	r.HandleFunc("/", indexGetHandler).Methods("GET")

	// r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	r.HandleFunc("/dash", dashGetHandler).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
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
