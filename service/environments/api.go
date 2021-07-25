package service

import (
	"github.com/badhouseplants/envspotting-go-proto/models/external/gitlab/environments"
	"google.golang.org/grpc"
)

type serviceGrpcImpl struct {
	environments.UnimplementedEnvironmentsServer
}

func Register(grpcServer *grpc.Server) {
	environments.RegisterEnvironmentsServer(grpcServer, NewapiGrpcImpl())
}

func NewapiGrpcImpl() *serviceGrpcImpl {
	return &serviceGrpcImpl{}
}
