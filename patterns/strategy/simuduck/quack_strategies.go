package simuduck

import "log"

type QuackStrategy interface {
	Quack()
}

type Quack struct{}

func (q *Quack) Quack() {
	log.Println("QUAAACK!")
}

type Squeak struct{}

func (s *Squeak) Quack() {
	log.Println("sss-qqq-uuuuuuu-eeeeeee-aaa-k")
}

type MuteDuck struct{}

func (md *MuteDuck) Quack() {
	log.Println("....[muted]...")
}
