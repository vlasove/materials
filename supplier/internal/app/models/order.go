package models

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/vlasove/materials/supplier/internal/app/helpers"
)

// Order ...
type Order struct {
	OrderUID          string  `json:"order_uid"`
	Entry             string  `json:"entry"`
	InternalSignature string  `json:"internal_signature"`
	Payment           Payment `json:"payment"`
	Items             []Items `json:"items"`
	Locale            string  `json:"locale"`
	CustomerID        string  `json:"customer_id"`
	TrackNumber       string  `json:"track_number"`
	DeliveryService   string  `json:"delivery_service"`
	Shardkey          string  `json:"shardkey"`
	SmID              int     `json:"sm_id"`
}

// FillRandomOrder ...
func (o *Order) FillRandomOrder() {
	o.OrderUID = uuid.New().String()
	o.Entry = helpers.RandStringRunes()
	o.InternalSignature = helpers.RandStringRunes()
	o.Payment = RandomPayment()
	o.Items = RandomItemsSlice()
	o.Locale = helpers.RandStringRunes()
	o.CustomerID = helpers.RandStringRunes()
	o.TrackNumber = helpers.RandStringRunes()
	o.DeliveryService = helpers.RandStringRunes()
	o.Shardkey = helpers.RandStringRunes()
	o.SmID = rand.Intn(100)
}
