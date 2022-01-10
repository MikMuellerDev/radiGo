module github.com/MikMuellerDev/radiGo/routes

go 1.17

require (
	github.com/MikMuellerDev/radiGo/audio v0.0.0-00010101000000-000000000000
	github.com/MikMuellerDev/radiGo/middleware v0.0.0-00010101000000-000000000000
	github.com/MikMuellerDev/radiGo/sessions v0.0.0-00010101000000-000000000000
	github.com/MikMuellerDev/radiGo/templates v0.0.0-00010101000000-000000000000
	github.com/MikMuellerDev/radiGo/utils v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace (
	github.com/MikMuellerDev/radiGo/audio => ../audio
	github.com/MikMuellerDev/radiGo/middleware => ../middleware
	github.com/MikMuellerDev/radiGo/routes => ../routes
	github.com/MikMuellerDev/radiGo/sessions => ../sessions
	github.com/MikMuellerDev/radiGo/templates => ../templates
	github.com/MikMuellerDev/radiGo/utils => ../utils
)
