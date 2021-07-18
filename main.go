package main

import (
	"fmt"
	"net"

	environments "github.com/badhouseplants/envspotting-gitlab/service/environments"
	projects "github.com/badhouseplants/envspotting-gitlab/service/projects"
	"github.com/badhouseplants/envspotting-gitlab/tools/logger"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

var (
	host string
	log  = logger.GetSimpleLogger()
)

func getHost() string {
	host = fmt.Sprintf("%s:%s", viper.GetString("app_host"), viper.GetString("app_port"))
	return host
}

func registerServices(grpcServer *grpc.Server) {
	environments.Register(grpcServer)
	projects.Register(grpcServer)
	// Disable on prod env
	reflection.Register(grpcServer)
}

func setupGrpcUnaryOpts() grpc.ServerOption {
	return grpc_middleware.WithUnaryServerChain(
		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_logrus.UnaryServerInterceptor(logger.GrpcLogrusEntry, logger.GrpcLogrusOpts...),
	)
}

func setupGrpcStreamOpts() grpc.ServerOption {
	return grpc_middleware.WithStreamServerChain(
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
		grpc_logrus.StreamServerInterceptor(logger.GrpcLogrusEntry, logger.GrpcLogrusOpts...),
	)
}


func main() {
	setDefaultVars()
	// seting up grpc server
	listener, err := net.Listen("tcp", getHost())
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(
		setupGrpcStreamOpts(),
		setupGrpcUnaryOpts(),
	)
	registerServices(grpcServer)
	log.Infof("starting to serve on %s", getHost())
	grpcServer.Serve(listener)
}

// set defaults values for viper variables
func setDefaultVars() {
	// app variables
	viper.SetDefault("app_host", "0.0.0.0")
	viper.SetDefault("app_port", "9090")
}
