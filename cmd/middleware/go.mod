module github.com/MikMuellerDev/radiGo/midddleware

go 1.17

replace github.com/MikMuellerDev/radiGo/sessions => ../sessions

require github.com/MikMuellerDev/radiGo/sessions v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
)
