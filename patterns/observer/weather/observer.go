package weather

import "log"

type Observer interface {
	Update(Data)
	GetName() string
	Display()
}

const currentConditions = "CurrentConditions"

type CurrentConditions struct {
	data    Data
	subject Subject
	name    string
}

func NewCurrentConditions(subject Subject) *CurrentConditions {
	cc := &CurrentConditions{
		subject: subject,
		name:    currentConditions,
	}
	subject.RegisterObserver(cc)
	return cc
}

func (cc *CurrentConditions) Update(data Data) {
	cc.data = data
	log.Println("current conditions was updated successfully")
}

func (cc *CurrentConditions) GetName() string {
	log.Println("current conditions was requesting for data")
	return cc.name
}

func (cc *CurrentConditions) Display() {
	log.Println("=====Current Conditions======")
	log.Println("		Temperature:", cc.data.Temperature)
	log.Println("		Pressure:", cc.data.Pressure)
	log.Println("=============================")
}

const generalDisplay = "GeneralDisplay"

type GeneralDisplay struct {
	data    Data
	subject Subject
	name    string
}

func NewGeneralDisplay(subject Subject) *GeneralDisplay {
	gd := &GeneralDisplay{
		subject: subject,
		name:    generalDisplay,
	}
	subject.RegisterObserver(gd)
	return gd
}

func (gd *GeneralDisplay) Update(data Data) {
	gd.data = data
	log.Println("general display was updated successfully")
}

func (gd *GeneralDisplay) GetName() string {
	log.Println("general display was requested for data")
	return gd.name
}

func (gd *GeneralDisplay) Display() {
	log.Println("=======General Display=====")
	log.Printf("		TemperatureGD: %f", gd.data.Temperature)
	log.Printf("		PressureGD: %f", gd.data.Pressure)
	log.Printf("		HumidityGD: %f", gd.data.Humidity)
	log.Println("===========================")
}
