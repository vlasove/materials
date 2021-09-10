package models

import (
	"math/rand"

	"github.com/vlasove/materials/test_src/generator/internal/app/helpers"
)

// Payment ...
type Payment struct {
	Transaction  string `json:"transaction"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
}

// RandomPayment ...
func RandomPayment() Payment {
	p := Payment{}
	p.Transaction = helpers.RandStringRunes()
	p.Currency = helpers.RandStringRunes()
	p.Provider = helpers.RandStringRunes()
	p.Amount = rand.Intn(1000)
	p.PaymentDt = rand.Intn(10000)
	p.Bank = helpers.RandStringRunes()
	p.DeliveryCost = rand.Intn(1000)
	p.GoodsTotal = rand.Intn(1000)
	return p
}

// // Validate ...
// func (u *User) Validate() error {
// 	return validation.ValidateStruct(
// 		u,
// 		validation.Field(&u.Email, validation.Required, is.Email),
// 		validation.Field(&u.Password, validation.By(RequiredIf(u.EncryptedPassword == "")), validation.Length(4, 100)),
// 	)
// }
