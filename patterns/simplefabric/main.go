package main

import "github.com/vlasove/materials/patterns/simplefabric/pizzashop"

const (
	cheesePizzaType    = "cheese"
	pepperoniPizzaType = "pepperoni"
	veggiePizzaType    = "veggie"
)

func main() {
	factory := pizzashop.NewPizzaFactory()
	pizzaStore := pizzashop.NewPizzaStore(factory)
	pizzaStore.OrderPizza(cheesePizzaType)
	pizzaStore.OrderPizza(pepperoniPizzaType)
	pizzaStore.OrderPizza(veggiePizzaType)
}
