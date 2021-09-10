package publisher

type Config struct {
	ClusterURLs []string `toml:"ClusterURLs"`
	ClusterID   string   `toml:"ClusterID"`
	ClientID    string   `toml:"ClientID"`
	Subject     string   `toml:"Subject"`
	Name        string   `toml:"Name"`
}

func NewConfig() *Config {
	return &Config{
		ClusterURLs: []string{"demo.nats.io"},
		ClusterID:   "test-cluster",
		Subject:     "test",
		Name:        "NATS API publisher",
		ClientID:    "test-client",
	}
}
