package git

import (
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"
)

type GitLab struct {
	Config *config.Config
}

type GitLabPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		URL string `json:"url"`
	} `json:"repository"`
}

func (h *GitLab) GitSync(body []byte, header http.Header) {
	signature := header.Get("X-Gitlab-Token")
	if !(subtle.ConstantTimeCompare([]byte(signature), []byte(h.Config.Secret)) == 1) {
		log.Println("Error: verification of the request secret didn't pass")
		return
	}

	var p GitLabPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling git webhook payload: %s", err)
		return
	}

	branch := strings.Split(p.Ref, "/")
	p.Ref = branch[len(branch)-1]
	if p.Ref != h.Config.GitBranch {
		log.Printf("Not calling notify, received update refers to %s, not %s", p.Ref, h.Config.GitBranch)
		return
	}
	change := GitChange{Kind: "git", Source: Source{URL: p.Repository.URL, Branch: p.Ref}}
	SendFluxNotification(&change)
}
