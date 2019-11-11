package image

import (
	"net/http"
)

type Handler interface {
	ImageSync(body []byte, header http.Header)
}
