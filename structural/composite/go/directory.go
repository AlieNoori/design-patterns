package main

import "fmt"

type Directory struct {
	components []Component
	name       string
}

func (d *Directory) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, d.name)
	for _, component := range d.components {
		component.search(keyword)
	}
}

func (d *Directory) add(c Component) {
	d.components = append(d.components, c)
}
