package client

// Config ...
type Config struct {
	NatsUrls    string `toml:"nats_urls"`
	NatsSubject string `toml:"nats_subject"`
	NatsName    string `toml:"nats_name"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		NatsUrls:    "nats://demo.nats.io:4222",
		NatsSubject: "go.test",
		NatsName:    "NATS API Subscriber",
	}
}
