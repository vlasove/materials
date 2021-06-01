package main

import "fmt"

type Fahrenheit float64
type Celcius float64

// FromFahrenheitToCelcius...
func FromFahrenheitToCelcius(fahr Fahrenheit) Celcius {
	res := (fahr - 32) * 5 / 9
	return Celcius(res)
}

type Foot float64
type Meter float64

// FromFootToMeter ...
func FromFootToMeter(foot Foot) Meter {
	return Meter(foot * 0.3048)
}

func main() {
	var (
		fahr Fahrenheit = 111
		foot Foot       = 111
	)
	fmt.Printf("%.2f Fahrenheit is %.2f Celcius\n", fahr, FromFahrenheitToCelcius(fahr))
	fmt.Printf("%.2f Foot is %.2f Meters\n", foot, FromFootToMeter(foot))
}
