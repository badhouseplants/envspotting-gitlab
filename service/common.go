package service

import (
	"context"
	"fmt"

	grpcusers "github.com/badhouseplants/envspotting-gitlab/internal/grpc-users"
	"github.com/badhouseplants/envspotting-gitlab/tools/metadata"
	"github.com/badhouseplants/envspotting-go-proto/models/common"
)

func GetGitlabToken(ctx context.Context) (string, error) {
	ctx = metadata.MetadataInternalProxy(ctx)
	accID, err := grpcusers.AuthorizationClient.ParseIdFromToken(ctx, &common.EmptyMessage{})
	if err != nil {
		fmt.Println("asdfasdfasdf")
		return "", err
	}

	account, err := grpcusers.AccountClient.SelfGet(ctx, accID)
	if err != nil {
		return "", err
	}

	token := account.GetGitlabToken()
	return token, nil
}
