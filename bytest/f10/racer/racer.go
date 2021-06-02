package racer

import (
	"fmt"
	"net/http"
	"time"
)

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

var tenSecondsTimeout = 10 * time.Second

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case val := <-ping(a):
		fmt.Println(val)
		return a, nil
	case val := <-ping(b):
		fmt.Println(val)
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
