package service

import (
	context "context"

	g "github.com/allanger/gitlab-environment-aggregator/third_party/gitlab"
	commonserv "github.com/badhouseplants/envspotting-gitlab/service"

	"github.com/allanger/gitlab-environment-aggregator/tools/logger"
	"github.com/badhouseplants/envspotting-go-proto/models/external/gitlab/environments"
	"github.com/xanzy/go-gitlab"
)

func (s *serviceGrpcImpl) Get(ctx context.Context, in *environments.EnvironmentID) (*environments.EnvironmentInfo, error) {
	log := logger.GetGrpcLogger(ctx)

	token, err := commonserv.GetGitlabToken(ctx)
	if err != nil {
		return nil, err
	}

	git, err := g.Client(token)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	env, _, err := git.Environments.GetEnvironment(int(in.ProjectId), int(in.Id))
	if err != nil {
		log.Error(err)
	}
	envOut := &environments.EnvironmentInfo{
		Id:    int64(env.ID),
		Name:  env.Name,
		State: env.State,
		Url:   env.ExternalURL,
	}
	if env.LastDeployment != nil {
		envOut = &environments.EnvironmentInfo{
			Ref:       env.LastDeployment.Ref,
			Sha:       env.LastDeployment.SHA,
			CiStatus:  env.LastDeployment.Status,
			CiId:      int64(env.LastDeployment.ID),
			UserId:    int64(env.LastDeployment.User.ID),
			UserName:  env.LastDeployment.User.Name,
			UpdatedAt: env.LastDeployment.UpdatedAt.String(),
		}
	}
	return envOut, nil
}

func (s *serviceGrpcImpl) List(in *environments.EnvironmentName, stream environments.Environments_ListServer) error {
	log := logger.GetGrpcLogger(stream.Context())

	token, err := commonserv.GetGitlabToken(stream.Context())
	if err != nil {
		return err
	}

	git, err := g.Client(token)
	if err != nil {
		log.Error(err)
		return err
	}

	listopt := &gitlab.ListOptions{PerPage: 5}
	opt := &gitlab.ListEnvironmentsOptions{Search: &in.Name, ListOptions: *listopt}
	gitEnvironments, _, err := git.Environments.ListEnvironments(int(in.ProjectId), opt)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, env := range gitEnvironments {
		envOut := &environments.EnvironmentInfo{
			Id:    int64(env.ID),
			Name:  env.Name,
			State: env.State,
			Url:   env.ExternalURL,
		}
		if err := stream.Send(envOut); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil

}
