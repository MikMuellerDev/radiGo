package sessions

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger
var Store *sessions.CookieStore

func InitLogger(logger *logrus.Logger) {
	log = logger
}

func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	log.Trace(fmt.Sprintf("Generated random seed for sessions: %s", string(b)))
	return string(b)
}

func Init(useRandomSeed bool) {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
		"1234567890"
	if useRandomSeed {
		Store = sessions.NewCookieStore([]byte(StringWithCharset(40, charset)))
	} else {
		// By using a static sting like "", no login is required when restarting the server
		// The session encryption key is static, cookies stay valid
		// If a logout is forced during development, turn on production mode
		log.Warn("\x1b[33mUsing a static string for session encryption. This is a security risk and should not be used in production.")
		Store = sessions.NewCookieStore([]byte(""))
	}
}
