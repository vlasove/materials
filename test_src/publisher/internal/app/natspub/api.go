package natspub

import (
	"log"

	"github.com/nats-io/nats.go"
)

// APINatsConnector ...
type APINatsConnector struct {
	config *NatsConfig
	nc     *nats.Conn
}

// New ...
func New(config *NatsConfig) *APINatsConnector {
	return &APINatsConnector{config: config}
}

// BuildConnection ...
func (a *APINatsConnector) BuildConnection() error {
	opts := []nats.Option{nats.Name(a.config.NatsName)}
	nc, err := nats.Connect(a.config.NatsUrls, opts...)
	if err != nil {
		return err
	}
	a.nc = nc
	return nil
}

// CloseConnection ...
func (a *APINatsConnector) CloseConnection() {
	a.nc.Close()
}

// GetConnect ...
func (a *APINatsConnector) GetConnect() *nats.Conn {
	return a.nc
}

// Publish ...
func (a *APINatsConnector) Publish(msg []byte) error {

	if err := a.nc.Publish(a.config.NatsSubject, msg); err != nil {
		return err
	}
	a.nc.Flush()
	if err := a.nc.LastError(); err != nil {
		return err
	}
	log.Printf("Published to [%s]\n", a.config.NatsSubject)
	return nil
}
