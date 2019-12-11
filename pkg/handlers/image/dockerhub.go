package image

import (
	"encoding/json"
	"log"
	"net/http"

	flux_api "github.com/fluxcd/flux/pkg/api/v9"
	"github.com/fluxcd/flux/pkg/image"
	"github.com/proskehy/flux-webhook-receiver/pkg/utils"
)

type DockerHub struct{}

type DockerHubPayload struct {
	Repository struct {
		Name string `json:"repo_name"`
	} `json:"repository"`
}

func (h *DockerHub) ImageSync(body []byte, header http.Header) {
	var p DockerHubPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling DockerHub webhook payload: %s", err)
		return
	}

	c := flux_api.Change{
		Kind: flux_api.ImageChange,
		Source: flux_api.ImageUpdate{
			Name: image.Name{
				Domain: "",
				Image:  p.Repository.Name,
			},
		},
	}

	utils.SendFluxNotification(&c)
}
