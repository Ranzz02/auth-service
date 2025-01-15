run: build
	@./bin/api.exe

migrate:
	@go run cmd/db/migrate/main.go

drop:
	@go run cmd/db/drop/main.go

build:
	@go build -o bin/api.exe cmd/api/main.go