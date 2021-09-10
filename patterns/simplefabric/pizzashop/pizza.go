package pizzashop

import "log"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type BoldPizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func NewPizza(name, dough, sauce string, toppings []string) Pizza {
	return &BoldPizza{
		name:     name,
		dough:    dough,
		sauce:    sauce,
		toppings: toppings,
	}
}

func (bp *BoldPizza) Prepare() {
	log.Println("preparing ", bp.GetName())
	log.Println("Tossing dough...")
	log.Println("Adding sauce...")
	log.Println("Adding toppings: ")
	for _, topping := range bp.toppings {
		log.Println("	", topping)
	}
}
func (bp *BoldPizza) Bake() {
	log.Println("Bake for 25 minutes at 350")
}
func (bp *BoldPizza) Cut() {
	log.Println("Cut pizza to diagonal slices")
}
func (bp *BoldPizza) Box() {
	log.Println("Place pizza in official PizzaStore box")
}
func (bp *BoldPizza) GetName() string {
	return bp.name
}
