package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is returned with there isnt enough balance in the wallet
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient balance")

// Bitcoin type
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit the given amount in the account
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns a balance in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw detucts the given amount from the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
