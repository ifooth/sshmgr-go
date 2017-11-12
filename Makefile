VERSION := 0.1

ROOTDIR := $(shell pwd)
PROJECTNAME := sshmgr-go

.PHONY: build
build:
	go build -o bin/ssh ssh.go

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 go build -o bin/ssh-linux-amd64-$(VERSION) ssh.go

.PHONY: darwin
darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/ssh-darwin-amd64-$(VERSION) ssh.go