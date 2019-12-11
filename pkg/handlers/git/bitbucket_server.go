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

type BitbucketServer struct {
	Config *config.Config
}

type BitbucketServerPayload struct {
	Changes []struct {
		RefID string `json:"refId"`
	} `json:"changes"`
	Repository struct {
		Links struct {
			Clone []struct {
				Href string `json:"href"`
				Name string `json:"name"`
			} `json:"clone"`
		} `json:"links"`
	} `json:"repository"`
}

func (h *BitbucketServer) GitSync(body []byte, header http.Header) {
	signature := header.Get("X-Hub-Signature")
	if len(h.Config.Secret) != 0 {
		valid := utils.VerifySignatureSHA256(signature, h.Config.Secret, body)
		if !valid {
			log.Printf("Error: verification of the request secret didn't pass")
			return
		}
	}

	var p BitbucketServerPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling git webhook payload: %s", err)
		return
	}

	var url, branch string
	for _, l := range p.Repository.Links.Clone {
		if l.Name == "ssh" {
			url = l.Href
		}
	}
	if len(p.Changes) > 0 {
		b := strings.Split(p.Changes[0].RefID, "/")
		branch = b[len(b)-1]
	}
	if branch != h.Config.GitBranch {
		log.Printf("Not calling notify, received update refers to %s, not %s", branch, h.Config.GitBranch)
		return
	}
	c := flux_api.Change{
		Kind: flux_api.GitChange,
		Source: flux_api.GitUpdate{
			URL:    url,
			Branch: branch,
		},
	}
	utils.SendFluxNotification(&c)
}
