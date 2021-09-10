package models

import (
	"math/rand"

	"github.com/vlasove/materials/supplier/internal/app/helpers"
)

// Items ...
type Items struct {
	ChrtID     int    `json:"chrt_id"`
	Price      int    `json:"price"`
	Rid        string `json:"rid"`
	Name       string `json:"name"`
	Sale       int    `json:"sale"`
	Size       string `json:"size"`
	TotalPrice int    `json:"total_price"`
	NmID       int    `json:"nm_id"`
	Brand      string `json:"brand"`
}

// RandomItems ...
func RandomItems() Items {
	i := Items{}
	i.ChrtID = rand.Intn(10000)
	i.Price = rand.Intn(10000)
	i.Rid = helpers.RandStringRunes()
	i.Name = helpers.RandStringRunes()
	i.Sale = rand.Intn(10)
	i.Size = helpers.RandStringRunes()
	i.TotalPrice = rand.Intn(1000)
	i.NmID = rand.Intn(100)
	i.Brand = helpers.RandStringRunes()
	return i
}

// RandomItemsSlice ...
func RandomItemsSlice() []Items {

	n := rand.Intn(12) + 3
	slice := make([]Items, n)
	for i := 0; i < n; i++ {
		slice[i] = RandomItems()
	}
	return slice
}
