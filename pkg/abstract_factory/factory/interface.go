package factory

import "github.com/sirupsen/logrus"

type car = interface {
	Forward()
	Back()
	SomeAction1()
}

type moto = interface {
	Forward()
	Back()
	SomeAction2()
}

type carCreator func(model string, speed int, logger *logrus.Logger) car
type motoCreator func(model string, speed int, logger *logrus.Logger) moto

// Factory is common factory interface
type Factory interface {
	// CreateCar returns new car created by factory
	CreateCar(model string, speed int) car
	// CreateMoto returns new moto created by factory
	CreateMoto(model string, speed int) moto
}
