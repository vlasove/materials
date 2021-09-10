package pizzashop

type PizzaStore struct {
	factory *SimplePizzaFactory
}

func NewPizzaStore(factory *SimplePizzaFactory) *PizzaStore {
	return &PizzaStore{
		factory: factory,
	}
}

func (ps *PizzaStore) OrderPizza(pizzaType string) Pizza {
	var pizza Pizza = ps.factory.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}
