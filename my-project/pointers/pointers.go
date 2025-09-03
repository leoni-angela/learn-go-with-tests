package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
// in Go is symbol is lowercase, then it is private
// outside the package it's defined in
// in this case to manipulate the balance, only via methods
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
	// pointers to structs are automatically dereferenced
	// so we don't need to do this:
	// (*w).balance += amount
}

func (w *Wallet) Balance() Bitcoin { 
	return w.balance 
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error{

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}