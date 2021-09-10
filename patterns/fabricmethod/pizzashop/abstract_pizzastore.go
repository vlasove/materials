package pizzashop

type abstractPizzaStore struct{}

func newAbstractPizzaStore() *abstractPizzaStore {
	return &abstractPizzaStore{}
}

func (aps *abstractPizzaStore) OrderPizza(pizzaType string) Pizza {
	var pizza Pizza = aps.createPizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

func (aps *abstractPizzaStore) createPizza(pizzaType string) Pizza {
	return newAbstractPizza("default pizza", "x", "xx", []string{})
}
