package image

import (
	"encoding/json"
	"log"
	"net/http"

	flux_api "github.com/fluxcd/flux/pkg/api/v9"
	"github.com/fluxcd/flux/pkg/image"
	"github.com/proskehy/flux-webhook-receiver/pkg/utils"
	"github.com/spf13/viper"
)

type Nexus struct{}

type NexusPayload struct {
	Component struct {
		Name string `json:"name"`
	} `json:"component"`
}

func (h *Nexus) ImageSync(body []byte, header http.Header) {
	cfgSecret := viper.GetString("DOCKER_WEBHOOK_SECRET")

	signature := header.Get("X-Nexus-Webhook-Signature")
	if len(cfgSecret) != 0 {
		valid := utils.VerifySignatureSHA1(signature, cfgSecret, body)
		if !valid {
			log.Printf("Error: verification of the request secret didn't pass")
			return
		}
	}

	var p NexusPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling Nexus webhook payload: %s", err)
		return
	}

	c := flux_api.Change{
		Kind: flux_api.ImageChange,
		Source: flux_api.ImageUpdate{
			Name: image.Name{
				Domain: "",
				Image:  p.Component.Name,
			},
		},
	}

	utils.SendFluxNotification(&c)
}
