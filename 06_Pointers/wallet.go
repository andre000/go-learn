package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Error! Trying to withdraw more than the current balance")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

/*
	Go copies values when you pass them to functions/methods
	so if you're writing a function that needs to mutate state
	you'll need it to take a pointer (*) to the thing you want to change.
*/
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
