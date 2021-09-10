package starbuss

type HouseBlend struct {
	description string
	price       float64
}

func NewHouseBlend() Coffee {
	return &HouseBlend{
		description: "HouseBlend",
		price:       0.89,
	}
}

func (h *HouseBlend) GetDescription() string {
	return h.description
}

func (h *HouseBlend) Cost() float64 {
	return h.price
}
