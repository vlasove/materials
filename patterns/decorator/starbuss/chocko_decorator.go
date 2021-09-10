package starbuss

type Chockolate struct {
	coffee   Coffee
	baseCost float64
}

func NewChockolate(coffee Coffee) Coffee {
	return &Chockolate{
		coffee:   coffee,
		baseCost: 0.2,
	}
}

func (c *Chockolate) GetDescription() string {
	return c.coffee.GetDescription() + ", Chockolate"
}

func (c *Chockolate) Cost() float64 {
	return c.coffee.Cost() + c.baseCost
}
