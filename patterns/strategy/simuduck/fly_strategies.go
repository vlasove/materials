package simuduck

import "log"

type FlyStrategy interface {
	Fly()
}

type FlyWithWings struct{}

func (fww *FlyWithWings) Fly() {
	log.Println("just flying with wings")
}

type FlyNoWay struct{}

func (fnw *FlyNoWay) Fly() {
	log.Println("can not flying :(")
}

type FlyWithRocket struct{}

func (fwr *FlyWithRocket) Fly() {
	log.Println("fly with rocket!!!! WHOOOSHH!!")
}
