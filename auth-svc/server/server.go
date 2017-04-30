package server

import (
	"net"
	"os"

	"github.com/inconshreveable/log15"

	"google.golang.org/grpc"
)

var logger = log15.New("context", "Authorization server")

// Run executes all setup needed and starts the server
// It takes a grpc server instance and executes it after setup completes
func Run(server *grpc.Server) {
	listener, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		logger.Error("failed to listen", "method", "Run", "message", err)
	}

	logger.Info("Listening on port "+os.Getenv("PORT"), "method", "Run")
	server.Serve(listener)
}
