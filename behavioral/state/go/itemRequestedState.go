package main

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d more\n", i.vendingMachine.itemPrice-money)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
