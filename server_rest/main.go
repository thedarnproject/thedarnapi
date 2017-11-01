package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"fmt"

	"github.com/Sirupsen/logrus"
	gw "github.com/thedarnproject/thedarnapi/api"
	"github.com/thedarnproject/thedarnapi/constants"
	"github.com/thedarnproject/thedarnapi/util"
)

func run() error {

	GRPCPort := util.GetEnvVarOrDefault(constants.GRPCPortEnvVarName, constants.GRPCListenPort)
	logrus.Infof("detected gRPC on port %v", GRPCPort)

	RESTPort := util.GetEnvVarOrDefault(constants.RESTPortEnvVarName, constants.RESTListenPort)

	echoEndpoint := flag.String("echo_endpoint", fmt.Sprintf("localhost:%v", GRPCPort), "gRPC endpoint")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterErrorInHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	logrus.Infof("starting REST server on port %v", RESTPort)
	return http.ListenAndServe(fmt.Sprintf(":%v", RESTPort), mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
