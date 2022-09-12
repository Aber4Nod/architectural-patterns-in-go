package prototype

import "fmt"

type folder struct {
	children []Inode
	name     string
}

func (f *folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.Print(indentation + indentation)
	}
}

func (f *folder) Clone() Inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func NewFolder(name string, children ...Inode) Inode {
	return &folder{
		name:     name,
		children: children,
	}
}
