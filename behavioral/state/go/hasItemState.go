package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) insertMoney(m int) error {
	return fmt.Errorf("Please select item first.")
}

func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first.")
}
