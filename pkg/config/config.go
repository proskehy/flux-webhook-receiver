package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	GitHost       string
	GitBranches   []string
	GitSecret     string
	DockerHost    string
	DockerSecret  string
	GitEnabled    bool
	DockerEnabled bool
)

func InitializeConfig() {
	viper.SetDefault("GIT_WEBHOOK_SECRET", "")
	viper.SetDefault("GIT_HOST", "github")
	viper.SetDefault("FLUX_DOCKER_HOST", "dockerhub")
	viper.SetDefault("DOCKER_WEBHOOK_SECRET", "")
	viper.SetDefault("GIT_BRANCHES", "master")
	viper.SetDefault("GIT_ENABLED", true)
	viper.SetDefault("DOCKER_ENABLED", true)
	viper.AutomaticEnv()
}

func PrintConfig() {
	if viper.GetBool("GIT_ENABLED") && viper.GetBool("DOCKER_ENABLED") {
		log.Printf("Current config: Git secret: %s, Git branches: %s, Git host: %s, Docker host: %s, Docker secret: %s",
			viper.GetString("GIT_WEBHOOK_SECRET"),
			viper.GetStringSlice("GIT_BRANCHES"),
			viper.GetString("GIT_HOST"),
			viper.GetString("FLUX_DOCKER_HOST"),
			viper.GetString("DOCKER_WEBHOOK_SECRET"),
		)
	} else if viper.GetBool("GIT_ENABLED") {
		log.Printf("Current config: Git secret: %s, Git branches: %s, Git host: %s, Docker webhooks disabled",
			viper.GetString("GIT_WEBHOOK_SECRET"),
			viper.GetStringSlice("GIT_BRANCHES"),
			viper.GetString("GIT_HOST"),
		)
	} else if viper.GetBool("DOCKER_ENABLED") {
		log.Printf("Current config: Git webhooks disabled, Docker host: %s, Docker secret: %s",
			viper.GetString("FLUX_DOCKER_HOST"),
			viper.GetString("DOCKER_WEBHOOK_SECRET"),
		)
	} else {
		log.Printf("Current config: Git and Docker webhooks disabled")
	}
}
