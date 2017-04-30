package controllers

import (
	"fmt"

	"github.com/inconshreveable/log15"

	"github.com/codesword/microservices-blog-api/pb/authorization"
	"google.golang.org/grpc"
)

var logger = log15.New("context", "controllers")

//Server implements the grpc service interface for the controller
type Server struct{}

// SetupServerInteface prepares the grpc server interaface for this service
func Setup() *grpc.Server {
	s := grpc.NewServer()
	svc := &Server{}
	authorization.RegisterAuthorizationSvcServer(s, svc)
	return s
}

func errorMessage(msg string) error { return fmt.Errorf("Error: %s", msg) }
