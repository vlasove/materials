package pizzashop

type CheesePizza struct {
	Pizza
}

func NewCheesePizza() *CheesePizza {
	bold := NewPizza("Cheese Pizza", "12.5", "cheese sauce", []string{"mozarella", "parmezan"})
	return &CheesePizza{bold}
}
