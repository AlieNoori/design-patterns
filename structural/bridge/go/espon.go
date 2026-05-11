package main

import "fmt"

type Espon struct{}

func (e *Espon) PrintFile() {
	fmt.Println("Printing by a Espon Printer")
}
