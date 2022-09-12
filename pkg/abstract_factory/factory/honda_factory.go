package factory

import "github.com/sirupsen/logrus"

type hondaFactory struct {
	carCreator func(model string, speed int, logger *logrus.Logger) car
	motoCreator func(model string, speed int, logger *logrus.Logger) moto
	logger     *logrus.Logger
}

// CreateCar creates new car
func (f *hondaFactory) CreateCar(model string, speed int) car {
	f.logger.Infof("hondaFactory creates car, model: '%s' ", model)
	return f.carCreator(model, speed, f.logger)
}

func (f *hondaFactory) CreateMoto(model string, speed int) moto {
	f.logger.Infof("hondaFactory creates moto, model: '%s' ", model)
	return f.motoCreator(model, speed, f.logger)
}

// NewHondaFactory creates new car factory
func NewHondaFactory(carCreator carCreator, motoCreator motoCreator, logger *logrus.Logger) Factory {
	return &hondaFactory{
		carCreator: carCreator,
		motoCreator: motoCreator,
		logger:     logger,
	}
}
