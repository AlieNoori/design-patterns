package main

import "fmt"

type Directory struct {
	children []Inode
	name     string
}

func (d *Directory) print(indentation string) {
	fmt.Println(indentation + d.name)
	for _, i := range d.children {
		i.print(indentation + indentation)
	}
}

func (d *Directory) clone() Inode {
	cloneDir := &Directory{name: d.name + "_clone", children: []Inode{}}

	for _, i := range d.children {
		clone := i.clone()
		cloneDir.children = append(cloneDir.children, clone)
	}

	return cloneDir
}
