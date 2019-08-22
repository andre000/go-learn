package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds shows when user tries to withdraw more than the current balance
var ErrInsufficientFunds = errors.New("Error! Trying to withdraw more than the current balance")

// Bitcoin int type
type Bitcoin int

// String Adds 'BTC' when converted to string
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct with balance
type Wallet struct {
	balance Bitcoin
}

/*
	Go copies values when you pass them to functions/methods
	so if you're writing a function that needs to mutate state
	you'll need it to take a pointer (*) to the thing you want to change.
*/

// Deposit adds given amount to the balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw removes given amount of the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance returns the current balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
