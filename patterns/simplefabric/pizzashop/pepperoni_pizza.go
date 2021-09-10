package pizzashop

import "log"

type PepperoniPizza struct {
	Pizza
}

func NewPepperoniPizza() *PepperoniPizza {
	bold := NewPizza("Pepperoni pizza", "15.5", "hot sauce", []string{"dople pepperoni slices"})
	return &PepperoniPizza{bold}
}

func (pp *PepperoniPizza) Cut() {
	log.Println("Cut pizza to square slices")
}
