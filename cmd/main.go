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
		if len(c.Secret) > 0 {
			log.Println("Warning: running bitbucket with secret set, bitbucket doesn't support secrets")
		}
		s.GitHandler = &git.Bitbucket{Config: c}
	case "bitbucket_server":
		s.GitHandler = &git.BitbucketServer{Config: c}
	default:
		c.GitHost = "github"
		c.GitBranch = "master"
		c.DockerHost = "dockerhub"
		s.ImageHandler = &image.DockerHub{}
		s.GitHandler = &git.GitHub{Config: c}
	}
	log.Printf("Running with config: secret: %s, git branch: %s, git host: %s, docker host: %s", c.Secret, c.GitBranch, c.GitHost, c.DockerHost)
	return s
}

func createConfig() *config.Config {
	s := os.Getenv("GIT_WEBHOOK_SECRET")
	gb := os.Getenv("GIT_BRANCH")
	gh := strings.ToLower(os.Getenv("GIT_HOST"))
	dh := strings.ToLower(os.Getenv("DOCKER_HOST"))
	return config.NewConfig(gh, gb, s, dh)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := initServer()
	http.HandleFunc("/gitSync", s.GitSync)
	http.HandleFunc("/imageSync", s.ImageSync)
	log.Fatal(http.ListenAndServe(":3033", nil))
}
