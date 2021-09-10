package pizzashop

type NyStyleCheesePizza struct {
	*abstractPizza
}

func NewNyStyleCheesePizza() *NyStyleCheesePizza {
	abs := newAbstractPizza(
		"NY Style Sauce and Cheese Pizza",
		"Thin Crust Dough",
		"Marinara Sauce",
		[]string{"Grated Reggiano Cheese"},
	)
	return &NyStyleCheesePizza{abs}
}

type NyStyleVeggiePizza struct {
	*abstractPizza
}

func NewNyStyleVeggiePizza() *NyStyleVeggiePizza {
	abs := newAbstractPizza(
		"NY Style Sauce and Veggie Pizza",
		"Thin Crust Dough",
		"Marinara Suace",
		[]string{"Grated Veggie Cheese"},
	)
	return &NyStyleVeggiePizza{abs}
}
