package config

type Config struct {
	GitHost      string
	GitBranch    string
	GitSecret    string
	DockerHost   string
	DockerSecret string
}

func NewConfig(gh, gb, gs, dh, ds string) *Config {
	return &Config{GitHost: gh, GitBranch: gb, GitSecret: gs, DockerHost: dh, DockerSecret: ds}
}
