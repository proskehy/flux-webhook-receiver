package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"log"
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
