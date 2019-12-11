package git

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	flux_api "github.com/fluxcd/flux/pkg/api/v9"
	"github.com/proskehy/flux-webhook-receiver/pkg/config"
	"github.com/proskehy/flux-webhook-receiver/pkg/utils"
)

type Bitbucket struct {
	Config *config.Config
}

type BitbucketPayload struct {
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Push struct {
		Changes []struct {
			New struct {
				Name string `json:"name"`
			} `json:"new"`
		} `json:"changes"`
	} `json:"push"`
}

func (h *Bitbucket) GitSync(body []byte, header http.Header) {
	// can't verify signature (bitbucket doesn't offer that functionality)

	var p BitbucketPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling git webhook payload: %s", err)
		return
	}

	var url = fmt.Sprintf("git@bitbucket.org:%s.git", p.Repository.FullName)
	var branch string
	if len(p.Push.Changes) > 0 {
		branch = p.Push.Changes[0].New.Name
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
