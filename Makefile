build:
	go build -o dps ./cmd/

test:
	go vet ./...
	go test ./... -v
