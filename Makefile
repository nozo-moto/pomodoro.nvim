NAME            := pomodoro.nvim
VERSION         := v0.0.1
GOPATH          ?= $(shell go env GOPATH)
XDG_CONFIG_HOME ?= $(shell echo $XDG_CONFIG_HOME)

default:
	make install 
	make clean

build:
	vgo build -ldflags "-w -s" -o bin/$(NAME)

install:
	make build
	cp bin/$(NAME) $(GOPATH)/bin/

clean:
	rm -rf bin/* vendor/*

docker-build:
	docker build -t nozomi0966/circleci_golang1.11_vgo .


