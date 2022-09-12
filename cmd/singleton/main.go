package main

import (
	"fmt"

	"github.com/aber4nod/architectural-patterns-in-go/pkg/singleton"
)

func main() {
	s := singleton.NewSingleton()
	for i := 0; i < 30; i++ {
		go s.GetInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
