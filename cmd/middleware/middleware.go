package middleware

import (
	"fmt"
	"net/http"

	sessions "github.com/MikMuellerDev/radiGo/sessions"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(logger *logrus.Logger) {
	log = logger
}

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		value, ok := session.Values["valid"]
		valid, okParse := value.(bool)

		query := r.URL.Query()
		username := query.Get("username")
		password := query.Get("password")

		if ok && okParse && valid {
			handler.ServeHTTP(w, r)
			return
		} else if TestCredentials(username, password) {

			// Saves session
			session, _ := sessions.Store.Get(r, "session")
			session.Values["valid"] = true
			session.Values["username"] = username
			session.Save(r, w)
			handler.ServeHTTP(w, r)
			return
		}
		log.Trace(fmt.Sprintf("Invalid Session, redirecting %s to /login", r.URL.Path))
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func LogRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Trace(fmt.Sprintf("[\x1b[32m%s\x1b[0m] FROM: (\x1b[34m%s\x1b[0m) [%s] Serving path:\x1b[35m%s\x1b[0m, user agent:%s", r.Method, r.RemoteAddr, r.Proto, r.URL.Path, r.UserAgent()))
		handler.ServeHTTP(w, r)
	}
}
