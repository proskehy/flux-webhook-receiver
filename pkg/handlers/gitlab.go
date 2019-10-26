package handlers

import (
	"crypto/subtle"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type GitLab struct {
	GitHost string
	Secret  string
}

type GitLabPayload struct {
	Ref        string `json:"ref"`
	Repository struct {
		URL string `json:"git_http_url"`
	} `json:"repository"`
}

func NewGitLab(gh string) *GitLab {
	return &GitLab{
		GitHost: gh,
	}
}

func (h *GitLab) GitSync(body []byte, header http.Header) {
	signature := header.Get("X-Gitlab-Token")
	if !(subtle.ConstantTimeCompare([]byte(signature), []byte(h.Secret)) == 1) {
		log.Println("Error: verification of the request secret didn't pass")
		return
	}

	var p GitLabPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling payload: %s", err)
		return
	}

	branch := strings.Split(p.Ref, "/")
	p.Ref = branch[len(branch)-1]
	log.Printf("Call localhost:3030/notify with payload %s", GitChange{Kind: "git", Source: Source{URL: p.Repository.URL, Branch: p.Ref}})
}
