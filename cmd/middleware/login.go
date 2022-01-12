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
	for _, v := range users {
		fmt.Println(v)
		if v.Name == user && v.Password == password {
			return true
		}
	}
	
	return false
}
