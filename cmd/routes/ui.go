package routes

import (
	"net/http"

	"github.com/MikMuellerDev/radiGo/middleware"
	"github.com/MikMuellerDev/radiGo/sessions"
	"github.com/MikMuellerDev/radiGo/templates"
)

// UI Pages
func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dash", http.StatusFound)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	value, ok := session.Values["valid"]
	valid, okParse := value.(bool)

	if ok && okParse && valid {
		http.Redirect(w, r, "/dash", http.StatusFound)
		return
	}

	templates.ExecuteTemplate(w, "login.html", http.StatusOK)
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

	if middleware.TestCredentials(username, password) {
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
