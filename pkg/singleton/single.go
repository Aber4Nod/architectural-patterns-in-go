package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single interface {
	GetInstance() Single
}

type single struct {
	once sync.Once
	instance Single
}

func (s *single) GetInstance() Single {
	if s.instance == nil {
		s.once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				s.instance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return s
}

func NewSingleton() Single {
	return &single{}
}