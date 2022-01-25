appname := radiGo-1.1.5
radiGoDir := radiGo

sources := $(wildcard *.go)

build =  cd ./cmd && GOOS=$(1) GOARCH=$(2) go build -o ../bin/$(appname)$(3)
tar =  mkdir -p build && cd ../ && tar -cvzf ./$(appname)_$(1)_$(2).tar.gz $(radiGoDir)/bin $(radiGoDir)/config $(radiGoDir)/static $(radiGoDir)/templates && mv $(appname)_$(1)_$(2).tar.gz $(radiGoDir)/build

.PHONY: all linux

all:	linux


##### LINUX BUILDS #####
linux: build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_386.tar.gz build/linux_amd64.tar.gz

build/linux_386.tar.gz: $(sources)
	$(call build,linux,386,)
	$(call tar,linux,386)

build/linux_amd64.tar.gz: $(sources)
	$(call build,linux,amd64,)
	$(call tar,linux,amd64)

build/linux_arm.tar.gz: $(sources)
	$(call build,linux,arm,)
	$(call tar,linux,arm)

build/linux_arm64.tar.gz: $(sources)
	$(call build,linux,arm64,)
	$(call tar,linux,arm64)

