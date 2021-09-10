package main

import "github.com/vlasove/materials/patterns/observer/weather"

func main() {
	originData := weather.NewData(36.6, 1024, 155)

	weatherDataServer := weather.NewWeatherData()
	_ = weather.NewCurrentConditions(weatherDataServer)

	weatherDataServer.SetData(originData)
	weatherDataServer.NotifyObservers()
	weatherDataServer.FlushAll()

	newData := weather.NewData(22, 22, 22)
	gd := weather.NewGeneralDisplay(weatherDataServer)
	weatherDataServer.SetData(newData)
	weatherDataServer.NotifyObservers()
	weatherDataServer.FlushAll()

	weatherDataServer.RemoveObserver(gd)
	weatherDataServer.FlushAll()
}
