package main

import (
	"fmt"
	"log"

	"github.com/vlasove/materials/patterns/fabricmethod/pizzashop"
)

func main() {
	nyStore := pizzashop.NewNyPizzaStore()
	nyPizzaCheese := nyStore.OrderPizza("cheese")
	log.Println("Ethan ordered a", nyPizzaCheese.GetName())
	fmt.Println()
	nyPizzaVeggie := nyStore.OrderPizza("veggie")
	log.Println("Joel ordered a", nyPizzaVeggie.GetName())
}
