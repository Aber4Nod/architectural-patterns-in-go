package car

// Car is common interface of cars
type Car = interface {
	// Forward commands car to move forwards
	Forward()
	// Back commands car to move back
	Back()
	// SomeAction1 commands car to perform some action1
	SomeAction1()
}
