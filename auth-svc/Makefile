export DATABASE_URL = postgres://postgres:password@127.0.0.1:5432/authorization?sslmode=disable
export GO_ENV = test
export MigrationsDir = $(shell pwd)/db/migrations
export PG_SCHEMA = default
export PORT = :30001

all: test

run:
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o artifact .
