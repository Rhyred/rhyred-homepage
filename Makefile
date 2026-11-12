.PHONY: dev build run generate clean test fmt

# Development
dev: generate run

# Generate templ templates
generate:
	templ generate

# Build the application
build: generate
	go build -o bin/server.exe ./cmd/server

# Run the application
run:
	go run ./cmd/server

# Run tests
test:
	go test ./...

# Format code
fmt:
	gofmt -w .
	templ fmt .

# Clean build artifacts
clean:
	rm -rf bin/

