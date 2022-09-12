package main

import (
	"fmt"

	"github.com/aber4nod/architectural-patterns-in-go/pkg/prototype"
)

func main() {
	file1 := prototype.NewFile("File1")
	file2 := prototype.NewFile("File2")
	file3 := prototype.NewFile("File3")

	folder1 := prototype.NewFolder("Folder1", file1)

	folder2 := prototype.NewFolder("Folder2", folder1, file2, file3)

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
