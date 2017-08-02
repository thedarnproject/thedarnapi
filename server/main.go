package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"net"

	api "github.com/thedarnproject/thedarnapi/api"
	"google.golang.org/grpc"
)

type errorInput struct{}

func (*errorInput) Submit(ctx context.Context, data *api.Data) (*api.Success, error) {
	return &api.Success{Success: true}, nil
}

func main() {
	port := flag.Int("p", 8080, "port to listen to")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterErrorInServer(grpcServer, &errorInput{})
	grpcServer.Serve(lis)

}
