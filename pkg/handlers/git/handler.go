package git

import (
	"net/http"
)

type Handler interface {
	GitSync(body []byte, header http.Header)
}
