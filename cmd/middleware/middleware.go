package middleware

import (
	"net/http"

	sessions "github.com/MikMuellerDev/radiGo/sessions"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		value, ok := session.Values["valid"]
		valid, okParse := value.(bool)
		if ok && okParse && valid {
			handler.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
