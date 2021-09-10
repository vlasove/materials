package client

import (
	"log"

	"github.com/nats-io/nats.go"
)

var i = 0

// APINatsConnector ...
type APINatsConnector struct {
	config *Config
	nc     *nats.Conn
}

// NewNatsConnector ...
func NewAPINatsConnector(config *Config) *APINatsConnector {
	return &APINatsConnector{config: config}
}

// BuildConnection ...
func (a *APINatsConnector) BuildConnection() error {
	opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
	opts = a.setupOptions(opts)
	nc, err := nats.Connect(a.config.NatsUrls, opts...)
	if err != nil {
		return err
	}
	a.nc = nc
	return nil
}

// GetConnection ...
func (a *APINatsConnector) GetConnection() *nats.Conn {
	return a.nc
}

// CloseConnection ...
func (a *APINatsConnector) CloseConnection() {
	a.nc.Close()
}

// Subscribe ...
func (a *APINatsConnector) Subscribe(f func(*nats.Msg)) error {
	_, err := a.nc.Subscribe(a.config.NatsSubject, f)
	if err != nil {
		return err
	}
	a.nc.Flush()

	if err := a.nc.LastError(); err != nil {
		return err
	}
	return nil
}

// setupOptions ...
func (a *APINatsConnector) setupOptions(opts []nats.Option) []nats.Option {
	// totalWait := 10 * time.Minute
	// reconnectDelay := time.Second

	// opts = append(opts, nats.ReconnectWait(reconnectDelay))
	// opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	// opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
	// 	log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	// }))
	// opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
	// 	log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	// }))
	// opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
	// 	log.Fatalf("Exiting: %v", nc.LastError())
	// }))
	return opts
}

// PrintAdapter ...
func (a *APINatsConnector) PrintAdapter(msg *nats.Msg) {
	i += 1
	printMsg(msg, i)
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("\n\n[#%d] Received on [%s]: '%s'\n\n", i, m.Subject, string(m.Data))
}
