package main

import "github.com/vlasove/materials/patterns/strategy/simuduck"

func ducking(d *simuduck.Duck) {
	d.Display()
	d.PerformFly()
	d.PerformQuack()
	d.Swim()
}

func main() {
	mallardDuck := simuduck.New(
		"MallardDuck",
		&simuduck.FlyWithWings{},
		&simuduck.Quack{},
	)
	ducking(mallardDuck)

	mallardDuck.SetFlySrategy(&simuduck.FlyWithRocket{})
	mallardDuck.SetQuackStrategy(&simuduck.MuteDuck{})
	ducking(mallardDuck)
}
