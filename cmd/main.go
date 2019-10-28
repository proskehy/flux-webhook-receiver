package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"

	"github.com/proskehy/flux-webhook-receiver/pkg/handlers"
	"github.com/proskehy/flux-webhook-receiver/pkg/server"
)

func initServer() (*server.Server, *config.Config) {
	c := createConfig()
	switch c.GitHost {
	case "github":
		return server.NewServer(&handlers.GitHub{Config: c}), c
	case "gitlab":
		return server.NewServer(&handlers.GitLab{Config: c}), c
	case "bitbucket":
		return server.NewServer(&handlers.Bitbucket{Config: c}), c
	case "bitbucket_server":
		return server.NewServer(&handlers.BitbucketServer{Config: c}), c
	default:
		c.GitHost = "github"
		c.GitBranch = "master"
		return server.NewServer(&handlers.GitHub{Config: c}), c
	}
}

func createConfig() *config.Config {
	s := os.Getenv("GIT_WEBHOOK_SECRET")
	gb := os.Getenv("GIT_BRANCH")
	gh := strings.ToLower(os.Getenv("GIT_HOST"))
	return config.NewConfig(gh, gb, s)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s, c := initServer()
	log.Printf("Running with config: secret: %s, git branch: %s, git host: %s", c.Secret, c.GitBranch, c.GitHost)
	http.HandleFunc("/gitSync", s.GitSync)
	log.Fatal(http.ListenAndServe(":3031", nil))
}
