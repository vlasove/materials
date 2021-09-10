package starbuss

type DarkRoast struct {
	description string
	price       float64
}

func NewDarkRoast() Coffee {
	return &DarkRoast{
		description: "DarkRoast",
		price:       0.99,
	}
}

func (d *DarkRoast) GetDescription() string {
	return d.description
}

func (d *DarkRoast) Cost() float64 {
	return d.price
}
