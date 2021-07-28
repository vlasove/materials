package hospital

import "log"

// Конкретный обработчик - осмотр доктора
type Doctor struct {
	Next Department
}

func NewDoctor(next Department) *Doctor {
	return &Doctor{
		Next: next,
	}
}

func (d *Doctor) SetNext(next Department) {
	d.Next = next
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		log.Println("doctor checkup already done")
		d.Next.Execute(p)
		return
	}
	log.Println("doctor checking patient")
	p.DoctorCheckUpDone = true
	d.Next.Execute(p)
}
