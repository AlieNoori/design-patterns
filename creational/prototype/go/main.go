package main

import "fmt"

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	directory1 := &Directory{
		children: []Inode{file1},
		name:     "Directory 1",
	}

	directory2 := &Directory{
		children: []Inode{directory1, file2, file3},
		name:     "Directory 2",
	}

	fmt.Println("\nPrinting hierarchy for Directory2")
	directory2.print("  ")

	cloneDir := directory2.clone()
	fmt.Println("\nPrinting hierarchy for clone Directory")
	cloneDir.print("  ")
}
