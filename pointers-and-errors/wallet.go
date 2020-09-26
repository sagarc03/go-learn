package wallet

// Wallet struct
type Wallet struct {
	balance int
}

// Deposit the given amount in the account
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance returns a balance in the wallet
func (w *Wallet) Balance() int {
	return w.balance
}
