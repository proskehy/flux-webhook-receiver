package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/proskehy/flux-webhook-receiver/pkg/handlers"
	"github.com/proskehy/flux-webhook-receiver/pkg/server"
)

func initServer() server.Server {
	s := os.Getenv("GIT_WEBHOOK_SECRET")
	gh := strings.ToLower(os.Getenv("GIT_HOST"))
	log.Printf("Config: secret: %s, git host: %s", s, gh)
	switch gh {
	case "github":
		return server.NewServer(&handlers.GitHub{GitHost: gh, Secret: s})
	case "gitlab":
		return server.NewServer(&handlers.GitLab{GitHost: gh, Secret: s})
	default:
		return server.NewServer(&handlers.GitHub{GitHost: "github", Secret: s})
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := initServer()
	http.HandleFunc("/gitSync", s.GitSync)
	log.Fatal(http.ListenAndServe(":3031", nil))
}
