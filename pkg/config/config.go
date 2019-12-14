package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	GitHost      string
	GitBranch    string
	GitSecret    string
	DockerHost   string
	DockerSecret string
)

func InitializeConfig() {
	viper.SetDefault("GIT_WEBHOOK_SECRET", "")
	viper.SetDefault("GIT_BRANCH", "master")
	viper.SetDefault("GIT_HOST", "github")
	viper.SetDefault("FLUX_DOCKER_HOST", "dockerhub")
	viper.SetDefault("DOCKER_WEBHOOK_SECRET", "")
	viper.AutomaticEnv()
}

func PrintConfig() {
	log.Printf("Current config: Git secret: %s, Git branch: %s, Git host: %s, Docker host: %s, Docker secret: %s",
		viper.GetString("GIT_WEBHOOK_SECRET"),
		viper.GetString("GIT_BRANCH"),
		viper.GetString("GIT_HOST"),
		viper.GetString("FLUX_DOCKER_HOST"),
		viper.GetString("DOCKER_WEBHOOK_SECRET"),
	)
}
