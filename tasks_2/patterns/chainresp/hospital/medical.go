package hospital

import "log"

// Конкретный обработчик - кабинет медикаментов
type Medical struct {
	Next Department
}

func NewMedical(next Department) *Medical {
	return &Medical{
		Next: next,
	}
}

func (m *Medical) SetNext(next Department) {
	m.Next = next
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		log.Println("medicine already given to patient")
		m.Next.Execute(p)
		return
	}
	log.Println("medical giving medicine to patient")
	p.MedicineDone = true
	m.Next.Execute(p)
}
