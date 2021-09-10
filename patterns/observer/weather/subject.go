package weather

type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	SetData(Data)
	NotifyObservers()
	FlushAll()
}

type WeatherData struct {
	data      Data
	observers []Observer
	changed   bool
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (wd *WeatherData) RegisterObserver(observer Observer) {
	wd.observers = append(wd.observers, observer)
}

func (wd *WeatherData) RemoveObserver(observer Observer) {
	var idToDelete int
	for id, obs := range wd.observers {
		if isEqual(obs, observer) {
			idToDelete = id
			break
		}
	}
	wd.observers = append(wd.observers[:idToDelete], wd.observers[idToDelete+1:]...)
}

func (wd *WeatherData) NotifyObservers() {
	if wd.changed {
		for _, observer := range wd.observers {
			observer.Update(wd.data)
		}
	}
	wd.changed = false
}

func (wd *WeatherData) SetData(data Data) {
	wd.data = data
	wd.changed = true
}

func (wd *WeatherData) FlushAll() {
	for _, observer := range wd.observers {
		observer.Display()
	}
}

func isEqual(lhs, rhs Observer) bool {
	return lhs.GetName() == rhs.GetName()
}
