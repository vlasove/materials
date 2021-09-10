package starbuss

type Espresso struct {
	description string
	price       float64
}

func NewEspresso() Coffee {
	return &Espresso{
		description: "Espresso",
		price:       1.99,
	}
}

func (e *Espresso) GetDescription() string {
	return e.description
}

func (e *Espresso) Cost() float64 {
	return e.price
}
