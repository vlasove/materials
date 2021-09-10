package simuduck

import "log"

type Duck struct {
	Name          string
	flyStrategy   FlyStrategy
	quackStrategy QuackStrategy
}

func New(name string, flyStrategy FlyStrategy, quackStrategy QuackStrategy) *Duck {
	return &Duck{
		Name:          name,
		flyStrategy:   flyStrategy,
		quackStrategy: quackStrategy,
	}
}

func (d *Duck) SetFlySrategy(f FlyStrategy) {
	d.flyStrategy = f
}

func (d *Duck) SetQuackStrategy(q QuackStrategy) {
	d.quackStrategy = q
}

func (d *Duck) PerformFly() {
	d.flyStrategy.Fly()
}

func (d *Duck) PerformQuack() {
	d.quackStrategy.Quack()
}

func (d *Duck) Display() {
	log.Printf("Look at this %s !!!", d.Name)
}

func (d *Duck) Swim() {
	log.Printf("%s can swimming!", d.Name)
}
