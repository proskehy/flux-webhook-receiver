package config

type Config struct {
	GitHost    string
	GitBranch  string
	Secret     string
	DockerHost string
}

func NewConfig(gh, gb, s, dh string) *Config {
	return &Config{GitHost: gh, GitBranch: gb, Secret: s, DockerHost: dh}
}
