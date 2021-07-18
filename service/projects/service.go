package service

import (
	"context"

	as "github.com/allanger/gitlab-environment-aggregator/service/users/accounts"
	g "github.com/allanger/gitlab-environment-aggregator/third_party/gitlab"
	"github.com/allanger/gitlab-environment-aggregator/tools/logger"
	"github.com/badhouseplants/envspotting-gitlab/models/external/gitlab/projects"
	"github.com/xanzy/go-gitlab"
)

func (s *serviceGrpcImpl) Get(ctx context.Context, in *projects.ProjectID) (*projects.ProjectInfo, error) {
	log := logger.GetGrpcLogger(ctx)
	token, err := as.GetGitlabTokenByID(ctx)
	git, err := g.Client(token)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	opt := &gitlab.GetProjectOptions{}
	project, _, err := git.Projects.GetProject(int(in.Id), opt)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	out := &projects.ProjectInfo{
		Id:          int64(project.ID),
		Name:        project.Name,
		WebUrl:      project.WebURL,
		AvatarUrl:   project.AvatarURL,
		Description: project.Description,
		ReadmeUrl:   project.ReadmeURL,
	}
	return out, nil
}

func (s *serviceGrpcImpl) List(in *projects.ProjectName, stream projects.Projects_ListServer) error {
	log := logger.GetGrpcLogger(stream.Context())
	token, err := as.GetGitlabTokenByID(stream.Context())
	git, err := g.Client(token)

	if err != nil {
		log.Error(err)
		return err
	}
	listopt := &gitlab.ListOptions{PerPage: 5}
	opt := &gitlab.ListProjectsOptions{Search: &in.Name, ListOptions: *listopt}
	gitProjects, _, err := git.Projects.ListProjects(opt)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, project := range gitProjects {
		prjectOut := &projects.ProjectInfo{
			Id:          int64(project.ID),
			Name:        project.Name,
			Description: project.Description,
			WebUrl:      project.WebURL,
			AvatarUrl:   project.AvatarURL,
			ReadmeUrl:   project.ReadmeURL,
		}
		if err := stream.Send(prjectOut); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil

}
