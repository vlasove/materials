package main

import (
	"fmt"

	"github.com/vlasove/materials/patterns/decorator/starbuss"
)

func main() {
	espresso := starbuss.NewEspresso()
	fmt.Printf("%s : $%f\n", espresso.GetDescription(), espresso.Cost())

	darkRoast := starbuss.NewDarkRoast()
	darkRoast = starbuss.NewWhip(darkRoast)
	darkRoast = starbuss.NewWhip(darkRoast)
	darkRoast = starbuss.NewChockolate(darkRoast)
	fmt.Printf("%s : $%f\n", darkRoast.GetDescription(), darkRoast.Cost())
}
