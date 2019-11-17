package git

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"
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
		valid := VerifySignatureSHA256(signature, h.Config.Secret, body)
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

	c := GitChange{Kind: "git"}
	for _, l := range p.Repository.Links.Clone {
		if l.Name == "ssh" {
			c.Source.URL = l.Href
		}
	}
	if len(p.Changes) > 0 {
		branch := strings.Split(p.Changes[0].RefID, "/")
		c.Source.Branch = branch[len(branch)-1]
	}
	if c.Source.Branch != h.Config.GitBranch {
		log.Printf("Not calling notify, received update refers to %s, not %s", c.Source.Branch, h.Config.GitBranch)
		return
	}
	SendFluxNotification(&c)
}
