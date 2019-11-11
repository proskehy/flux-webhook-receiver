package image

import (
	"encoding/json"
	"log"
	"net/http"
)

type Flux struct{}

type FluxPayload struct {
	Kind   string `json:"kind"`
	Source struct {
		Name struct {
			Domain string `json:"domain"`
			Image  string `json:"image"`
		} `json:"name"`
	} `json:"source"`
}

func (h *Flux) ImageSync(body []byte, header http.Header) {
	var p FluxPayload

	if err := json.Unmarshal(body, &p); err != nil {
		log.Printf("Error unmarshalling image webhook payload: %s", err)
		return
	}

	log.Printf("Call localhost:3030/notify with payload %s", p)
}
