run: build
	@./bin/api.exe

build:
	@go build -o bin/api.exe cmd/api/main.go