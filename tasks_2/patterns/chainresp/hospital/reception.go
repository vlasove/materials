package hospital

import "log"

// Конкретный обработчик - применое отделение
type Reception struct {
	Next Department
}

func NewReception(next Department) *Reception {
	return &Reception{
		Next: next,
	}
}

func (r *Reception) SetNext(next Department) {
	r.Next = next
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		log.Println("patient registration already done")
		r.Next.Execute(p)
		return
	}
	log.Println("reception registering patient")
	p.RegistrationDone = true
	r.Next.Execute(p)
}
