package main

import (
	"github.com/sirupsen/logrus"

	"github.com/aber4nod/architectural-patterns-in-go/pkg/abstract_factory/car"
	"github.com/aber4nod/architectural-patterns-in-go/pkg/abstract_factory/factory"
	"github.com/aber4nod/architectural-patterns-in-go/pkg/abstract_factory/moto"
)

const (
	sportCarModel = "Sport Car"
	sportCarSpeed = 500

	motoModel = "Moto"
	motoSpeed = 100
)

func main() {
	logger := logrus.New()

	bmwFactory := factory.NewBMWFactory(car.NewCar, moto.NewMoto, logger)
	bmwCar := bmwFactory.CreateCar(sportCarModel, sportCarSpeed)
	bmwCar.Forward()
	bmwCar.Back()

	hondaFactory := factory.NewHondaFactory(car.NewCar, moto.NewMoto, logger)
	hondaMoto := hondaFactory.CreateMoto(motoModel, motoSpeed)
	hondaMoto.Forward()
	hondaMoto.Back()
}
