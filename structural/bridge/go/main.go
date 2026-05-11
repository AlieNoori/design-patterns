package main

import "fmt"

func main() {
	hpPrinter := &Hp{}
	esponPrinter := &Espon{}

	macComputer := &Mac{}

	macComputer.setPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.setPrinter(esponPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(esponPrinter)
	winComputer.Print()
	fmt.Println()
}
