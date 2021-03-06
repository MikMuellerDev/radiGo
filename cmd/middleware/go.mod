module github.com/MikMuellerDev/radiGo/midddleware

go 1.17

replace (
	github.com/MikMuellerDev/radiGo/sessions => ../sessions
	github.com/MikMuellerDev/radiGo/utils => ../utils
)

require (
	github.com/MikMuellerDev/radiGo/sessions v0.0.0-00010101000000-000000000000
	github.com/MikMuellerDev/radiGo/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)
