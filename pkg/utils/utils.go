package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	flux_api "github.com/fluxcd/flux/pkg/api/v9"
)

func VerifySignatureSHA1(signature, secret string, payload []byte) bool {
	if len(signature) == 0 {
		log.Println("Error: request without signature")
		return false
	}

	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
		return false
	}
	return true
}

func VerifySignatureSHA256(signature, secret string, payload []byte) bool {
	if len(signature) == 0 {
		log.Println("Error: request without signature")
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(signature[7:]), []byte(expectedMAC)) {
		return false
	}
	return true
}

func SendFluxNotification(c *flux_api.Change) {
	requestBody, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error marshalling payload: %s", err)
	}

	log.Printf("Notifying Flux about %s change", c.Kind)

	resp, err := http.Post("http://localhost:3030/api/flux/v11/notify", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("Error delivering Flux notification: %s", err)
	}

	if resp != nil {
		resp.Body.Close()
	}
}
