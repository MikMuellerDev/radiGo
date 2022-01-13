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
		// TODO impl checker for api requests that use url params auth instead of session
		if ok && okParse && valid {
			log.Trace(fmt.Sprintf("Valid Session, serving %s", r.URL.Path))
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
