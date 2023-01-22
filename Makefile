DSN=postgres://postgres:150903@localhost:5432/greenlight?sslmode=disable

run:
	go run cmd/api/*.go

up:
	migrate -path migrations -database "${DSN}" -verbose up

down:
	migrate -path migrations -database "${DSN}" -verbose down

build:
	go build -o bin\main.exe cmd/api/

.PHONY: run up build