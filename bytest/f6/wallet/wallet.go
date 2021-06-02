// Package wallet ...
package wallet

import (
	"errors"
	"fmt"
)

var (
	// ErrInsufficientFunds  ...
	ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
)

// Bitcoin ...
type Bitcoin int

// String ...
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

// Deposit ...
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw ...
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
