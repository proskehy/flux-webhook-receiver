package git

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"
)

type GitHub struct {
	Config *config.Config
}

type GitHubPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		URL string `json:"ssh_url"`
	} `json:"repository"`
}

func (h *GitHub) GitSync(body []byte, header http.Header) {
	signature := header.Get("X-Hub-Signature")
	if len(h.Config.Secret) != 0 {
		valid := VerifySignatureSHA1(signature, h.Config.Secret, body)
		if !valid {
			log.Printf("Error: verification of the request secret didn't pass")
			return
		}
	}

	var p GitHubPayload

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
