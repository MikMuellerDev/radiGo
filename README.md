# RadiGo
## Version 1.0.0
Headless internet radio for linux

### What is it?
RadiGo allows you to turn your RaspberryPi or old PC into a headless internet-radio or remote-controlled speaker. The functionality of RadiGo is to send commands to that server using the sleek and modern webinterface.

### Why does it exist?
If your plan is to run a full media center than you might like [Jellyfin](https://github.com/jellyfin/). But if you are confused with setting up KODI on your server and using the webinterface than RadiGo might be just the right tool for the job.

I originally had the goal to create a DIY-internet radio that runs on my Raspberry Pi (Raspbian) and looked at KODI.
When it was time to add the internet-radio stations via KODI, I realized that I run Raspbian in headless mode which didn't allow me yo set up the stations.
Because staying headless (no display) was my goal, KODI was not a solution to my problem.

I also took a short look at [Volumio](https://volumio.com/) for the Raspberry Pi, but it is meant to be a standalone Operating System which din't meet my criteria either.

### Why you should use it
RadiGo is lightweight, offers streaming support for jellyfin and is compatible with almost every internet radio stream out there.
Due to RadiGo being written in Go, the application consist of a single binary for every common Linux architecture.
Note: for Jellyfin casting to work, install [Jellyfin MPV shim](https://github.com/jellyfin/jellyfin-mpv-shim/blob/master/README.md#linux-installation). If you are using it for the radio feature, install [MPV](https://mpv.io/manual/master/) on your host.
Other dependencies should be minimal, thanks to the binary, which contains all Go modules.

-> The Go toolchain / programming language is not required

### Features
- Sleek and modern webinterface
- Mobile friendly
- Use your server as a jellyfin cast client, for more information visit [this](https://github.com/jellyfin/jellyfin-mpv-shim#readme) website (Jellyfin MPV shim).
- Play **any** internet radio station on your server you like.


###

## Dashboard
![](./README_ASSETS/dashboard.png)
![](./README_ASSETS/login.png)
![](./README_ASSETS/terminal.png)