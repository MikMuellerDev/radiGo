appname := radiGo-1.3.1
radiGoDir := radiGo

sources := $(wildcard *.go)

build =cd ./cmd && GOOS=$(1) GOARCH=$(2) go build -o ../bin/$(appname)$(3) $(4)
tar =  mkdir -p build && cd ../ && tar --exclude $(radiGoDir)/static/js/src -cvzf ./$(appname)_$(1)_$(2).tar.gz $(radiGoDir)/bin $(radiGoDir)/config $(radiGoDir)/static $(radiGoDir)/templates && mv $(appname)_$(1)_$(2).tar.gz $(radiGoDir)/build

.PHONY: all linux

all:	linux


run: web
	cd ./cmd && go run .


web:
	tsc -b
	minify-all-js static/js/out

build: web all linux clean

clean:
	rm -rf bin
	rm -rf log
	rm -rf static/js/out

cleanall: clean
	rm -rf build

##### LINUX BUILDS #####
linux: build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_386.tar.gz build/linux_amd64.tar.gz

build/linux_386.tar.gz: $(sources)
	$(call build,linux,386,)
	$(call tar,linux,386)

build/linux_amd64.tar.gz: $(sources)
	$(call build,linux,amd64, -ldflags '-extldflags "-fno-PIC -static"' -buildmode pie -tags 'osusergo netgo static_build')
	$(call tar,linux,amd64)

build/linux_arm.tar.gz: $(sources)
	$(call build,linux,arm,)
	$(call tar,linux,arm)

build/linux_arm64.tar.gz: $(sources)
	$(call build,linux,arm64,)
	$(call tar,linux,arm64)

