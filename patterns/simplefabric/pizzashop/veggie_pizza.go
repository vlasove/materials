package pizzashop

import "log"

type VeggiePizza struct {
	Pizza
}

func NewVeggiePizza() *VeggiePizza {
	bold := NewPizza("Veggie Pizza", "20.0", "tomato", []string{"cucmbers"})
	return &VeggiePizza{bold}
}

func (vp *VeggiePizza) Cut() {
	log.Println("Cut pizza to circle slices")
}

func (vp *VeggiePizza) Box() {
	log.Println("Place pizza in official PizzaStore eco-friendly box")
}
