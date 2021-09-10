package pizzashop

type NyPizzaStore struct {
	*abstractPizzaStore
}

func NewNyPizzaStore() *NyPizzaStore {
	abs := newAbstractPizzaStore()
	return &NyPizzaStore{
		abstractPizzaStore: abs,
	}
}

func (nps *NyPizzaStore) createPizza(pizzaType string) Pizza {
	switch pizzaType {
	case "cheese":
		return NewNyStyleCheesePizza()
	case "veggie":
		return NewNyStyleVeggiePizza()
	default:
		return nil
	}
}
