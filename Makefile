.PHONY: client
client:
	@echo "Building client binary"
	go build -o bin/client cmd/client/main.go
