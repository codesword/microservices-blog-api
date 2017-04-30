package main

import (
	"github.com/codesword/microservices-blog-api/auth-svc/controllers"
	"github.com/codesword/microservices-blog-api/auth-svc/server"
)

func main() {
	grpcServer := controllers.Setup()
	server.Run(grpcServer)
}
