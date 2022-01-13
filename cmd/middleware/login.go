package middleware

import (
	"fmt"

	"github.com/MikMuellerDev/radiGo/utils"
)

var users []utils.User

func InitializeLogin(config *utils.Config) {
	users = config.Users
}

func TestCredentials(user string, password string) bool {
	log.Trace(fmt.Sprintf("Testing login credentials for User %s", user))
	for _, v := range users {
		if v.Name == user && v.Password == password {
			log.Debug(fmt.Sprintf("Login successful for User %s", user))
			return true
		}
	}
	log.Debug(fmt.Sprintf("Login failed for User %s with Password: %s", user, password))
	return false
}
