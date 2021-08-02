package clock

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const (
	// BaseHost константа для обращения к стандартному серверу времени
	BaseHost = "0.beevik-ntp.pool.ntp.org"
)

// Clock - базовые часы
type Clock struct {
	writer      io.Writer
	response    *ntp.Response
	timePrecise time.Time
	timeLocal   time.Time
	host        string
}

// New ...
func New(host string, errorWriter io.Writer) *Clock {
	return &Clock{
		host:   host,
		writer: errorWriter,
	}
}

func (c *Clock) hostChecker() error {
	response, err := ntp.Query(c.host)
	if err != nil {
		return err
	}
	c.response = response
	return nil
}

func (c *Clock) setCurrent() error {
	if err := c.hostChecker(); err != nil {
		return err
	}
	c.timePrecise = time.Now().Add(c.response.ClockOffset)
	c.timeLocal = time.Now()
	return nil
}

func (c *Clock) String() string {
	if err := c.setCurrent(); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Precise:%v\nLocal:%v", c.timePrecise, c.timeLocal)
}
