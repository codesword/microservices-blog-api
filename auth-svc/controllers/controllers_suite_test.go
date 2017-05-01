package controllers

import (
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/codesword/microservices-blog-api/auth-svc/server"
	"github.com/codesword/microservices-blog-api/pb/authorization"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"testing"
)

func TestControllers(t *testing.T) {
	grpcServer := grpc.NewServer()
	svc := &Server{}
	authorization.RegisterAuthorizationSvcServer(grpcServer, svc)
	go server.Run(grpcServer)
	time.Sleep(time.Second * 5)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var addr = fmt.Sprintf("0.0.0.0%s", os.Getenv("PORT"))
var (
	conn   *grpc.ClientConn
	err    error
	client authorization.AuthorizationSvcClient
)

var _ = BeforeSuite(func() {
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		Fail("did not connect ")
	}
	client = authorization.NewAuthorizationSvcClient(conn)
})

var _ = AfterSuite(func() {
	conn.Close()
})
