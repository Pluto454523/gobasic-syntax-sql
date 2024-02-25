package syntax

import "fmt"

type Wallet struct {
	Name  string
	Money int
}

func InterfaceTutorial2() {

	// ** Create slice of Struct Wallet
	Wallets := []Wallet{
		{Name: "Wallet 1", Money: 2000},
		{Name: "Wallet 2", Money: 1000},
	}
	fmt.Printf("Wallets -> %p - %v\n", Wallets, Wallets)

	// ** Pass by reference
	acc1 := &Wallets[0]
	acc1.Money = 5000
	fmt.Printf("\nacc1 	 -> %p - %v\n", acc1, acc1)
	fmt.Printf("Wallets -> %p - %v\n", Wallets, Wallets)

	// ** Pass by value
	acc2 := Wallets[1]
	acc2.Money = 9999
	fmt.Printf("\nacc2 	 -> %p - %v\n", &acc2, acc2)
	fmt.Printf("Wallets -> %p - %v\n", Wallets, Wallets)

	// ** Change value by ref in for-range
	// ** 	=> เนื่องจาก Wallets เก็บ ข้อมูลเป็น struct ซุึ่งเป็น value type
	// **	=> Wallets[i]
	for i := range Wallets {
		fmt.Printf("\nWallets -> %p - %v\n", Wallets, Wallets)
		Wallets[i].Money += 1000
		fmt.Printf("Wallets[%v] -> %p - %v\n", i, &Wallets[i], Wallets[i])
		fmt.Printf("Wallets -> %p - %v\n", Wallets, Wallets)
	}

	// ! Careful
	for _, acc := range Wallets {
		acc.Money += 1000
		fmt.Printf("\nacc(%v) -> %v\n", acc.Name, acc)
		fmt.Printf("Wallets -> %v\n", Wallets)
	}
}
