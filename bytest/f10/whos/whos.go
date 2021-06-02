package whos

import (
	"errors"
	"net/http"
	"time"
)

var (
	defaultTimeout = 10 * time.Second
	ErrTimeout     = errors.New("too long requesting, timeout occured")
)

func WhosFasterURL(urls ...string) (winner string, err error) {
	return ConfigurableWhosFasterURL(defaultTimeout, urls...)
}

func ConfigurableWhosFasterURL(timeout time.Duration, urls ...string) (winner string, err error) {
	select {
	case url := <-faster(urls...):
		return url, nil
	case <-time.After(timeout):
		return "", ErrTimeout
	}
}

func faster(urls ...string) chan string {
	ch := make(chan string)
	for _, url := range urls {
		go func(u string) {
			http.Get(u)
			ch <- u
		}(url)
	}
	return ch
}
