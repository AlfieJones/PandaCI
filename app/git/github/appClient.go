package gitGithub

import (
	"context"
	"net/http"
	"strconv"

	gitShared "github.com/alfiejones/panda-ci/app/git/shared"
	"github.com/alfiejones/panda-ci/pkg/utils/env"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v68/github"
)

// TODO - We should cache these clients to avoid having to regenerate access tokens
func (c *GithubClient) newInstallationClient(ctx context.Context, installationID string) (*GithubInstallationClient, error) {
	appID, err := env.GetGithubAppID()
	if err != nil {
		return nil, err
	}

	key, err := env.GetGithubAppPrivateKey()
	if err != nil {
		return nil, err
	}

	installationIDInt, err := strconv.Atoi(installationID)
	if err != nil {
		return nil, err
	}

	itr, err := ghinstallation.New(http.DefaultTransport, int64(*appID), int64(installationIDInt), []byte(*key))
	if err != nil {
		return nil, err
	}

	client := github.NewClient(&http.Client{Transport: itr})

	graphqlClient := &graphqlClient{
		itr: itr,
	}

	return &GithubInstallationClient{
		githubClient:   client,
		installationID: installationID,
		gitClient:      c,
		graphqlClient:  graphqlClient,
		itr:            itr,
	}, nil
}

func (c *GithubClient) NewInstallationClient(ctx context.Context, installationID string) (gitShared.InstallationClient, error) {
	return c.newInstallationClient(ctx, installationID)
}

func (c *GithubClient) NewAppClient(ctx context.Context) (gitShared.AppClient, error) {
	appID, err := env.GetGithubAppID()
	if err != nil {
		return nil, err
	}

	key, err := env.GetGithubAppPrivateKey()
	if err != nil {
		return nil, err
	}

	atr, err := ghinstallation.NewAppsTransport(http.DefaultTransport, int64(*appID), []byte(*key))
	if err != nil {
		return nil, err
	}

	client := github.NewClient(&http.Client{Transport: atr})

	return &GithubAppClient{
		githubClient: client,
		gitClient:    c,
	}, nil
}
