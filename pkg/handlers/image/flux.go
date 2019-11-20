package image

import (
	"bytes"
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

	requestBody, err := json.Marshal(p)
	if err != nil {
		log.Printf("Error marshalling payload: %s", err)
	}

	log.Printf("Notifying Flux about %s change", p.Kind)

	resp, err := http.Post("http://localhost:3030/api/flux/v11/notify", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("Error delivering Flux notification: %s", err)
	}

	if resp != nil {
		resp.Body.Close()
	}
}
