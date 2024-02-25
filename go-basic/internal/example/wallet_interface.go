package example

import (
	"errors"
)

type Customer struct {
	CustomerID int
	Firstname  string
	Lastname   string
	Accounts   *walletAccount
}

type WallletService interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type walletAccount struct {
	name    string
	number  string
	balance float64
}

func NewWalletAccount(name, number string, bal float64) *walletAccount {
	return &walletAccount{name: name, number: number, balance: bal}
}

func (a *walletAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

func (a *walletAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *walletAccount) Balance() float64 {
	return a.balance
}
