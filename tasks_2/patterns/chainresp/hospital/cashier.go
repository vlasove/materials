package hospital

import "log"

type Cashier struct {
	Next Department
}

func NewCashier(next Department) *Cashier {
	return &Cashier{
		Next: next,
	}
}

func (c *Cashier) SetNext(next Department) {
	c.Next = next
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		log.Println("payment done")
		return
	}
	log.Println("cashier getting money from patient patient")
}
