package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Count int
}

func (s *SpySleeper) Sleep() {
	s.Count++
}

const (
	sleep = "sleep"
	write = "write"
)

type CountdownOperationSpy struct {
	counts []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.counts = append(c.counts, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.counts = append(c.counts, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("check answer and num of calls to sleep", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
		if sleeper.Count != 4 {
			t.Errorf("not enough calls to sleep. got %d want %d", sleeper.Count, 4)
		}
	})

	t.Run("check right sleep/write sequence at calls", func(t *testing.T) {
		spy := &CountdownOperationSpy{}
		Countdown(spy, spy)
		want := []string{
			sleep, write,
			sleep, write,
			sleep, write,
			sleep, write,
		}
		if !reflect.DeepEqual(spy.counts, want) {
			t.Errorf("got %v want %v", spy.counts, want)
		}
	})

}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
