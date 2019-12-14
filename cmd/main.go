package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"

	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/git"
	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/image"
	"github.com/proskehy/flux-webhook-receiver/pkg/server"
)

func initServer() *server.Server {
	s := server.NewServer()
	c := createConfig()
	switch c.GitHost {
	case "github":
		s.GitHandler = &git.GitHub{Config: c}
	case "gitlab":
		s.GitHandler = &git.GitLab{Config: c}
	case "bitbucket":
		if len(c.GitSecret) > 0 {
			log.Println("Warning: running bitbucket with secret set, bitbucket doesn't support secrets")
		}
		s.GitHandler = &git.Bitbucket{Config: c}
	case "bitbucket_server":
		s.GitHandler = &git.BitbucketServer{Config: c}
	default:
		c.GitHost = "github"
		c.GitBranch = "master"
		s.GitHandler = &git.GitHub{Config: c}
	}
	switch c.DockerHost {
	case "dockerhub":
		if len(c.DockerSecret) > 0 {
			log.Println("Warning: running dockerhub with secret set, dockerhub doesn't support secrets")
		}
		s.ImageHandler = &image.DockerHub{}
	case "nexus":
		s.ImageHandler = &image.Nexus{Config: c}
	default:
		c.DockerHost = "dockerhub"
		s.ImageHandler = &image.DockerHub{}
	}
	log.Printf("Running with config: Git secret: %s, git branch: %s, git host: %s, docker host: %s, docker secret: %s", c.GitSecret, c.GitBranch, c.GitHost, c.DockerHost, c.DockerSecret)
	return s
}

func createConfig() *config.Config {
	gs := os.Getenv("GIT_WEBHOOK_SECRET")
	gb := os.Getenv("GIT_BRANCH")
	gh := strings.ToLower(os.Getenv("GIT_HOST"))
	dh := strings.ToLower(os.Getenv("DOCKER_HOST"))
	ds := strings.ToLower(os.Getenv("DOCKER_WEBHOOK_SECRET"))
	return config.NewConfig(gh, gb, gs, dh, ds)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := initServer()
	http.HandleFunc("/gitSync", s.GitSync)
	http.HandleFunc("/imageSync", s.ImageSync)
	log.Fatal(http.ListenAndServe(":3033", nil))
}
