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
shadow: vet
	shadow ./...
.PHONY:shadow
golangci: shadow
	golangci-lint run
.PHONY:golangci
build: golangci
	go build -o ./hello.exe ./main.go
.PHONY:build
