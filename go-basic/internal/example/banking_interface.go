package example

import (
	"errors"
	"fmt"
)

type AccountService interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type account struct {
	owner   string
	number  string
	balance float64
}

func NewAccount(owner, number string, initialBalance float64) *account {
	return &account{owner: owner, number: number, balance: initialBalance}
}

func (a *account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

func (a *account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if amount > a.balance {
		return errors.New("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func (a *account) Balance() float64 {
	return a.balance
}

func BankingExample() {
	// Create a new account
	acc := NewAccount("Nawin K", "30039", 500.0)

	// Deposit ทดสอบถอนเงินออกจากบัญชี
	err := acc.Deposit(100.0)
	if err != nil {
		fmt.Println("Error: Deposit error:", err)
	} else {
		fmt.Println("Logs: Deposit Success")
	}

	// Withdraw ทดสอบฝากเงินเข้าบัญชี
	err = acc.Withdraw(30.0)
	if err != nil {
		fmt.Println("Error: Withdrawal error:", err)
	} else {
		fmt.Println("Logs: Withdrawal Success")
	}

	// Current balance
	fmt.Println("Current balance:", acc.Balance())
}
