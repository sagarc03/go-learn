package wallet

// Bitcoin type
type Bitcoin int

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
