package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"
	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/git"
	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/image"
	"github.com/proskehy/flux-webhook-receiver/pkg/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := server.NewServer()
	config.InitializeConfig()
	config.PrintConfig()
	if viper.GetBool("GIT_ENABLED") {
		gh := git.GetGitHandler(viper.GetString("GIT_HOST"))
		s.GitHandler = gh
		http.HandleFunc("/gitSync", s.GitSync)
	}
	if viper.GetBool("DOCKER_ENABLED") {
		ih := image.GetImageHandler(viper.GetString("FLUX_DOCKER_HOST"))
		s.ImageHandler = ih
		http.HandleFunc("/imageSync", s.ImageSync)
	}
	log.Fatal(http.ListenAndServe(":3033", nil))
}
