package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 123)
	fmt.Println()

	if err := walletFacade.addMoneyToWallet("abc", 123, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()

	if err := walletFacade.deductMoneyFromWallet("abc", 123, 5); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
