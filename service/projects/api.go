package service

import (
	"github.com/badhouseplants/envspotting-go-proto/models/external/gitlab/projects"
	"google.golang.org/grpc"
)

type serviceGrpcImpl struct {
	projects.UnimplementedProjectsServer
}

func Register(grpcServer *grpc.Server) {
	projects.RegisterProjectsServer(grpcServer, NewapiGrpcImpl())
}

func NewapiGrpcImpl() *serviceGrpcImpl {
	return &serviceGrpcImpl{}
}
