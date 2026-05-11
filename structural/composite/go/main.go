package main

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	directory1 := &Directory{
		name: "directory1",
	}

	directory1.add(file1)

	directory2 := &Directory{
		name: "directory2",
	}
	directory2.add(file2)
	directory2.add(file3)
	directory2.add(directory1)

	directory2.search("rose")
}
