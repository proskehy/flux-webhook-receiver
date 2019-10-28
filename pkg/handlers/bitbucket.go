package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"
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
		log.Printf("Error unmarshalling payload: %s", err)
		return
	}

	c := GitChange{Kind: "git"}
	c.Source.URL = fmt.Sprintf("git@bitbucket.org:%s.git", p.Repository.FullName)
	if len(p.Push.Changes) > 0 {
		c.Source.Branch = p.Push.Changes[0].New.Name
	}
	if c.Source.Branch != h.Config.GitBranch {
		log.Printf("Not calling notify, received update refers to %s, not %s", c.Source.Branch, h.Config.GitBranch)
		return
	}
	log.Printf("Call localhost:3030/notify with payload %s", c)
}
