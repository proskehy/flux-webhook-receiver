package git

import (
	"net/http"
)

type Handler interface {
	GitSync(body []byte, header http.Header)
}

func GetGitHandler(handler string) Handler {
	switch handler {
	case "github":
		return &GitHub{}
	case "gitlab":
		return &GitLab{}
	case "bitbucket":
		return &Bitbucket{}
	case "bitbucket_server":
		return &BitbucketServer{}
	default:
		return &GitHub{}
	}
}
