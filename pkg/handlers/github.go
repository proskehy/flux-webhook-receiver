package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/proskehy/flux-webhook-receiver/pkg/utils"
)

type GitHub struct {
	GitHost string
	Secret  string
}

type GitHubPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		URL string `json:"url"`
	} `json:"repository"`
}

func NewGitHub(gh string) *GitHub {
	return &GitHub{
		GitHost: gh,
	}
}

func (h *GitHub) GitSync(body []byte, header http.Header) {
	signature := header.Get("X-Hub-Signature")
	if len(h.Secret) != 0 {
		valid := utils.VerifySignatureSHA1(signature, h.Secret, body)
		if !valid {
			log.Printf("Error: verification of the request secret didn't pass")
			return
		}
	}

	var p GitHubPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling payload: %s", err)
		return
	}

	branch := strings.Split(p.Ref, "/")
	p.Ref = branch[len(branch)-1]
	log.Printf("Call localhost:3030/notify with payload %s", GitChange{Kind: "git", Source: Source{URL: p.Repository.URL, Branch: p.Ref}})
}
