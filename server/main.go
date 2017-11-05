package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	api "github.com/thedarnproject/thedarnapi/api"
	"github.com/thedarnproject/thedarnapi/constants"
	"github.com/thedarnproject/thedarnapi/datasources/stackoverflow"
	"github.com/thedarnproject/thedarnapi/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type errorInput struct{}

func (*errorInput) Submit(ctx context.Context, data *api.Data) (*api.Fix, error) {
	logrus.Infof("received data:\n%#v", *data)
	if len(data.Plugin) == 0 || len(data.Platform) == 0 || len(data.Error) == 0 {
		return nil, fmt.Errorf("invalid data submitted, make sure all the fields are populated: %v", data)
	}
	fix, err := stackoverflow.DarnIt(data.Error)
	if err != nil {
		return nil, errors.Wrap(err, "unable to DarnIt for StackOverflow data source")
	}

	return &api.Fix{Fix: fix}, nil
}

func main() {

	listenPort, err := strconv.Atoi(util.GetEnvVarOrDefault(
		constants.GRPCPortEnvVarName,
		constants.GRPCListenPort))
	if err != nil {
		logrus.Fatalf("unable to convert string to int: %v", err)
	}

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
