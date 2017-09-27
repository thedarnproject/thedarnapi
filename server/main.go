package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Sirupsen/logrus"
	api "github.com/thedarnproject/thedarnapi/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	listenPort = 8081
)

type errorInput struct{}

func (*errorInput) Submit(ctx context.Context, data *api.Data) (*api.Success, error) {
	logrus.Infof("received data:\n%#v", *data)
	if len(data.Plugin) == 0 || len(data.Platform) == 0 || len(data.Error) == 0 {
		return nil, fmt.Errorf("invalid data submitted, make sure all the fields are populated: %v", data)
	}
	return &api.Success{Success: true}, nil
}

func main() {
	port := flag.Int("p", listenPort, "port to listen to")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("server failed to start listening: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterErrorInServer(grpcServer, &errorInput{})

	logrus.Infof("starting the server at %v", listenPort)
	if grpcServer.Serve(lis) != nil {
		log.Fatalf("error serving the API: %v", err)
	}
}
