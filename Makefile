build:
	go clean -cache
	go build -o ./bin/gin_wire

lint:
	goimports -w ./
	go vet ./...

test:
	go clean -cache -testcache
	go test ./...

wire:
	wire ./di/
