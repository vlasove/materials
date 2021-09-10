package pizzashop

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type PizzaStore interface {
	OrderPizza(string) Pizza
	createPizza(string) Pizza
}
