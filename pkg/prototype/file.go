package prototype

import "fmt"

type file struct {
	name string
}

func (f *file) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) Clone() Inode {
	return &file{name: f.name + "_clone"}
}

func NewFile(name string) Inode {
	return &file{
		name: name,
	}
}
