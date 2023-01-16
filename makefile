.PHONY: build

build:
	go build -v ./cmd/app

.PHONY: test-report test-unit test-integration	

test-report:
	go test -json ./... | go-test-report

test-unit:
	go test -v -run Unit -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-integration:
	go test -v -run Integration ./...

.DEFAULT_GOAL := build