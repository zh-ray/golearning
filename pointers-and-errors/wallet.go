package errors

import (
	"errors"
	"fmt"
)

// Bitcoin is a type
type Bitcoin int

// Wallet defines a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit calculate the deposit
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Println("address of balance in Deposit is ", &w.balance)
	w.balance += amount
}

// Balance return the balence
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFundsError represents the Error
var ErrInsufficientFundsError = errors.New("can't withdraw, insufficient funds")

//Withdraw calculate the withdraw
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFundsError
	}
	w.balance -= amount
	return nil
}

//String defines the format of Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
