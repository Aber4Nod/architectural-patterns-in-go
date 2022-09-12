package moto

// Moto is common interface of motos
type Moto = interface {
	// Forward commands moto to move forwards
	Forward()
	// Back commands moto to move back
	Back()
	// SomeAction2 commands moto to perform some action2
	SomeAction2()
}
