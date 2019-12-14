package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/proskehy/flux-webhook-receiver/pkg/config"
	"github.com/proskehy/flux-webhook-receiver/pkg/handlers"
	"github.com/proskehy/flux-webhook-receiver/pkg/server"
)

func createServer() *server.Server {
	s := server.NewServer()
	config.InitializeConfig()
	config.PrintConfig()
	hm := handlers.InitializeHandlerMap()
	gh := hm.GitHandlers[viper.GetString("GIT_HOST")]
	dh := hm.ImageHandlers[viper.GetString("FLUX_DOCKER_HOST")]
	s.GitHandler = gh
	s.ImageHandler = dh
	return s
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := createServer()
	http.HandleFunc("/gitSync", s.GitSync)
	http.HandleFunc("/imageSync", s.ImageSync)
	log.Fatal(http.ListenAndServe(":3033", nil))
}
