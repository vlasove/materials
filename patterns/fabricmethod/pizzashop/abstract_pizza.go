package pizzashop

import "log"

type abstractPizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func newAbstractPizza(name, dough, sauce string, toppings []string) *abstractPizza {
	return &abstractPizza{
		name:     name,
		dough:    dough,
		sauce:    sauce,
		toppings: toppings,
	}
}

func (ap *abstractPizza) Prepare() {
	log.Println("Preparing", ap.name)
	log.Println("Tossing dough....")
	log.Println("Adding sauce....")
	log.Println("Adding toppings:")
	for _, topping := range ap.toppings {
		log.Println("	", topping)
	}
}

func (ap *abstractPizza) Bake() {
	log.Println("Bake for 25 minutes at 350")
}

func (ap *abstractPizza) Cut() {
	log.Println("Cutting the pizza into diagonal slices")
}

func (ap *abstractPizza) Box() {
	log.Println("Place pizza in official PizzaStore box")
}

func (ap *abstractPizza) GetName() string {
	return ap.name
}
