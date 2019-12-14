package handlers

import (
	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/git"
	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/image"
)

type HandlerMap struct {
	GitHandlers   map[string]git.Handler
	ImageHandlers map[string]image.Handler
}

func InitializeHandlerMap() *HandlerMap {
	gh := map[string]git.Handler{
		"github":           &git.GitHub{},
		"gitlab":           &git.GitLab{},
		"bitbucket":        &git.Bitbucket{},
		"bitbucket_server": &git.BitbucketServer{},
	}

	ih := map[string]image.Handler{
		"dockerhub": &image.DockerHub{},
		"nexus":     &image.Nexus{},
	}

	return &HandlerMap{
		GitHandlers:   gh,
		ImageHandlers: ih,
	}
}
