package handlers

import (
	"net/http"
)

type Handler interface {
	GitSync(body []byte, header http.Header)
}
