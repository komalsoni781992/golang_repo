.DEFAULT_GOAL := build
fmt:
	go fmt ./...
.PHONY:fmt
lint: fmt
	golint ./...
.PHONY:lint
vet: fmt
	go vet ./...
.PHONY:vet
build: vet
	go build -o ./hello.exe ./hello_world/001_hello_world.go
.PHONY:build
