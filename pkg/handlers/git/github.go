package git

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	flux_api "github.com/fluxcd/flux/pkg/api/v9"
	"github.com/proskehy/flux-webhook-receiver/pkg/config"
	"github.com/proskehy/flux-webhook-receiver/pkg/utils"
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
	if len(h.Config.GitSecret) != 0 {
		valid := utils.VerifySignatureSHA1(signature[5:], h.Config.GitSecret, body)
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

	c := flux_api.Change{
		Kind: flux_api.GitChange,
		Source: flux_api.GitUpdate{
			URL:    p.Repository.URL,
			Branch: p.Ref,
		},
	}
	utils.SendFluxNotification(&c)
}
