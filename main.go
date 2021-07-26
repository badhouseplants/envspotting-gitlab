package main

import (
	"fmt"
	"net"

	grpcusers "github.com/badhouseplants/envspotting-gitlab/internal/grpc-users"
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
	log  = logger.GetServerLogger()
)

func init() {
	viper.SetDefault("envspotting_gitlab_host", "0.0.0.0")
	viper.SetDefault("envspotting_gitlab_port", "9090")
	viper.AutomaticEnv() // read in environment variables that match)
}

func getHost() string {
	host = fmt.Sprintf("%s:%s", viper.GetString("envspotting_gitlab_host"), viper.GetString("envspotting_gitlab_port"))
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
	// seting up grpc server
	listener, err := net.Listen("tcp", getHost())
	if err != nil {
		log.Fatal(err)
	}
	grpcusers.Connect()
	grpcServer := grpc.NewServer(
		setupGrpcStreamOpts(),
		setupGrpcUnaryOpts(),
	)
	registerServices(grpcServer)
	log.Infof("starting to serve on %s", getHost())
	grpcServer.Serve(listener)
}