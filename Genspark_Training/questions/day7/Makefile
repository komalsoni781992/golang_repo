.DEFAULT_GOAL := build
fmt:
	go fmt ./...
.PHONY:fmt
imports: fmt
	goimports -l -w .
.PHONY:imports
lint: imports
	golint ./...
.PHONY:lint
vet: lint
	go vet ./...
.PHONY:vet
golangci: vet
	golangci-lint run
.PHONY:golangci
build: golangci
	go build -o ./interface.exe ./main.go
.PHONY:build
