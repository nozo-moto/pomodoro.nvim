NAME            := pomodoro.nvim
VERSION         := v0.0.1
GOPATH          ?= $(shell go env GOPATH)
XDG_CONFIG_HOME ?= $(shell echo $XDG_CONFIG_HOME)

default:
	make deps
	make install 
	make clean

deps:
	dep ensure

build:
	go build -ldflags "-w -s" -o bin/$(NAME)

install:
	make build
	cp bin/$(NAME) $(GOPATH)/bin/

clean:
	rm -rf bin/* vendor/*


