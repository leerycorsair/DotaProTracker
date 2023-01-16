.PHONY: build

build:
	go build -v ./cmd/app

.PHONY: test-report test-unit test-integration	

test-unit:
	go test -v -run Unit -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-unit-report:
	go test -json -run Unit ./... | go-test-report

test-integration:
	go test -v -run Integration ./...

test-integration-report:
	go test -json -run Integration ./... | go-test-report

test-e2e:
	go test -v -run E2E ./...

test-e2e-report:
	go test -json -run E2E ./... | go-test-report

.DEFAULT_GOAL := build