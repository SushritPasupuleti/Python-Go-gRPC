run:
	@echo "Running server..."
	@go run main.go

build:
	@echo "Building server..."
	@GOOS=linux GOARCH=amd64 go build -o bin/server main.go 
	@GOOS=windows GOARCH=amd64 go build -o bin/server.exe main.go
	@GOOS=darwin GOARCH=amd64 go build -o bin/server.darwin main.go
	@echo "Done building for linux, windows and mac (x64 only)"
	@echo "Binaries are in bin/ directory"

run_build:
	@echo "Running server..."
	@./bin/server

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@golangci-lint run

dev:
	@air
