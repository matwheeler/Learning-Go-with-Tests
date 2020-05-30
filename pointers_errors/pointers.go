package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// The string function is used when call a %s when using fmt.Println() for example.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	return
}

var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}
