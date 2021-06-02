package main

import (
	"os"
	"time"

	"github.com/vlasove/materials/bytest/f8/mocking"
)

func main() {
	sleeper := mocking.NewSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
