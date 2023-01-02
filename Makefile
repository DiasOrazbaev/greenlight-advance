DSN=postgres://postgresql:password@localhost:1234/greenlight?sslmode=disable

run:
	go run cmd/api/*.go

up:
	migrate -path migrations -database "${DSN}" -verbose up

down:
	migrate -path migrations -database "${DSN}" -verbose down

build:
	go build -o bin cmd/api/*.go

.PHONY: run up build