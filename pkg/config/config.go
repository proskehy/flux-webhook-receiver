package config

type Config struct {
	GitHost   string
	GitBranch string
	Secret    string
}

func NewConfig(gh, gb, s string) *Config {
	return &Config{GitHost: gh, GitBranch: gb, Secret: s}
}
