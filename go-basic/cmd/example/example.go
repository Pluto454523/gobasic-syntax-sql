package main

import (
	"fmt"

	ex "github.com/pluto454523/basic-syntax/internal/example"
)

func main() {

	// ex.BankingExample()

	acc := ex.NewWalletAccount("Wallet A", "30039", 500.0)

	err := acc.Deposit(100.0)
	if err != nil {
		fmt.Println("Error: Deposit error:", err)
	} else {
		fmt.Println("Logs: Deposit Success")
	}

	err = acc.Withdraw(30.0)
	if err != nil {
		fmt.Println("Error: Withdrawal error:", err)
	} else {
		fmt.Println("Logs: Withdrawal Success")
	}

	fmt.Printf("Current balance: %v\n", acc.Balance())
	fmt.Printf("%p - %v\n", acc, acc)

	cust1 := ex.Customer{
		CustomerID: 1,
		Firstname:  "Nawin",
		Lastname:   "K",
		Accounts:   acc,
		// Accounts:   ex.NewWalletAccount("Wallet 1", "30039", 100.00),
	}

	fmt.Printf("Current balance: %v\n", cust1.Accounts.Balance())
	fmt.Printf("%p - %v\n", cust1.Accounts, cust1.Accounts)
}
