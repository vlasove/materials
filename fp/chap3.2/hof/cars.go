package hof

var CarsDB = initCarsDB()

func initCarsDB() []IndexedCar {
	indexedCars := []IndexedCar{}
	for i, car := range LoadCars() {
		indexedCars = append(indexedCars, IndexedCar{i, car})
	}
	lenCars := len(indexedCars)
	for i, car := range LoadMoreCars() {
		indexedCars = append(indexedCars, IndexedCar{i + lenCars, car})
	}
	return indexedCars
}

func LoadCars() Collection {
	return CsvToStruct("cars.csv")
}
