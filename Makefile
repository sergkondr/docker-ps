APP_NAME := docker-cps
APP_VERSION := v0.0.3

build: test
	go build -ldflags="-X 'main.version=${APP_VERSION}'" -o ${APP_NAME} ./cmd/

test:
	go vet ./...
	go test -v ./...