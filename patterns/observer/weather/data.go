package weather

type Data struct {
	Temperature float64
	Pressure    float64
	Humidity    float64
}

func NewData(temperature, pressure, humidity float64) Data {
	return Data{
		Temperature: temperature,
		Pressure:    pressure,
		Humidity:    humidity,
	}
}
