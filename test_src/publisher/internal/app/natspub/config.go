package natspub

// NatsConfig ...
type NatsConfig struct {
	NatsUrls    string `toml:"nats_urls"`
	NatsSubject string `toml:"nats_subject"`
	NatsName    string `toml:"nats_name"`
}

// NewNatsConfig ...
func NewNatsConfig() *NatsConfig {
	return &NatsConfig{
		NatsUrls:    "nats://demo.nats.io:4222",
		NatsSubject: "go.test",
		NatsName:    "NATS API publisher",
	}
}
