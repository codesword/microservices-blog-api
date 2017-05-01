export GO_ENV = test
export PORT = :30001

all: test

test:
	go test $$(go list ./... | grep -v /vendor/)
run:
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o artifact .
