package image

import (
	"net/http"
)

type Handler interface {
	ImageSync(body []byte, header http.Header)
}

func GetImageHandler(handler string) Handler {
	switch handler {
	case "dockerhub":
		return &DockerHub{}
	case "nexus":
		return &Nexus{}
	default:
		return &DockerHub{}
	}
}
