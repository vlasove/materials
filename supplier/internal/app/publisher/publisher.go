package publisher

import (
	"log"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/vlasove/materials/supplier/internal/app/models"
)

type Connector struct {
	NATSConn *nats.Conn
	Stream   stan.Conn
}

func NewConnector() *Connector {
	return &Connector{}
}

func Start(configPath string, sourcePath string) error {
	cfg := NewConfig()
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return err
	}
	conn := NewConnector()
	defer conn.Close()

	if err := conn.configureConnector(cfg); err != nil {
		return err
	}

	data, err := conn.getSourceData(sourcePath)
	if err != nil {
		return err
	}

	if err := conn.SendBatchWithPause(cfg, data, time.Second*2); err != nil {
		return err
	}

	return nil
}

func (c *Connector) configureConnector(cfg *Config) error {
	opts := []nats.Option{nats.Name(cfg.Name)}
	nc, err := nats.Connect(
		strings.Join(cfg.ClusterURLs, ","),
		opts...,
	)
	if err != nil {
		return err
	}
	c.NATSConn = nc

	sc, err := stan.Connect(cfg.ClusterID, cfg.ClientID, stan.NatsConn(nc))
	if err != nil {
		return err
	}
	c.Stream = sc
	log.Println("connection to NATS builded successfully")
	return nil

}

func (c *Connector) getSourceData(sourcePath string) ([]*models.Order, error) {
	or := NewOrderReader()
	if err := or.Open(sourcePath); err != nil {
		return nil, err
	}

	if err := or.Unmarshal(); err != nil {
		return nil, err
	}
	return or.orders, nil

}

func (c *Connector) SendBatchWithPause(cfg *Config, batch []*models.Order, pause time.Duration) error {
	for _, ord := range batch {
		time.Sleep(pause)
		if data, err := toByte(ord); err != nil {
			log.Println(err)
			continue
		} else {
			errStream := c.Stream.Publish(cfg.Subject, data)
			if err != nil {
				return errStream
			}
			log.Printf("Published [%s] : '%v'\n\n\n", cfg.Subject, *ord)
		}

	}
	return nil
}

func (c *Connector) Close() error {
	defer c.NATSConn.Close()
	if err := c.Stream.Close(); err != nil {
		return err
	}
	return nil

}
