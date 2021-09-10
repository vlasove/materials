package pizzashop

type SimplePizzaFactory struct{}

func NewPizzaFactory() *SimplePizzaFactory {
	return &SimplePizzaFactory{}
}

func (spf *SimplePizzaFactory) CreatePizza(pizzaType string) Pizza {
	var pizza Pizza
	switch pizzaType {
	case "cheese":
		pizza = NewCheesePizza()
	case "pepperoni":
		pizza = NewPepperoniPizza()
	case "veggie":
		pizza = NewVeggiePizza()
	}
	return pizza
}
