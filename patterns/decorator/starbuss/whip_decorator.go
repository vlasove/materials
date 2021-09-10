package starbuss

type Whip struct {
	coffee   Coffee
	baseCost float64
}

func NewWhip(coffee Coffee) Coffee {
	return &Whip{
		coffee:   coffee,
		baseCost: 0.1,
	}
}

func (w *Whip) GetDescription() string {
	return w.coffee.GetDescription() + ", Whip"
}

func (w *Whip) Cost() float64 {
	return w.coffee.Cost() + w.baseCost
}
