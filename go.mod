module github.com/MikMuellerDev/radiGo

go 1.17

require (
	github.com/MikMuellerDev/radiGo/routes v0.0.0-00010101000000-000000000000
	github.com/MikMuellerDev/radiGo/templates v0.0.0-00010101000000-000000000000
	github.com/MikMuellerDev/radiGo/utils v0.0.0-00010101000000-000000000000
)

replace (
	github.com/MikMuellerDev/radiGo/audio => ./cmd/audio
	github.com/MikMuellerDev/radiGo/middleware => ./cmd/middleware
	github.com/MikMuellerDev/radiGo/routes => ./cmd/routes
	github.com/MikMuellerDev/radiGo/sessions => ./cmd/sessions
	github.com/MikMuellerDev/radiGo/templates => ./cmd/templates
	github.com/MikMuellerDev/radiGo/utils => ./cmd/utils
)
