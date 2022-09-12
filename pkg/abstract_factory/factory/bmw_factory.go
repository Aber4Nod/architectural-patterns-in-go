package factory

import "github.com/sirupsen/logrus"

type bmwFactory struct {
	carCreator func(model string, speed int, logger *logrus.Logger) car
	motoCreator func(model string, speed int, logger *logrus.Logger) moto
	logger     *logrus.Logger
}

// CreateCar creates new car
func (f *bmwFactory) CreateCar(model string, speed int) car {
	f.logger.Infof("bmwFactory creates car, model: '%s' ", model)
	return f.carCreator(model, speed, f.logger)
}

func (f *bmwFactory) CreateMoto(model string, speed int) moto {
	f.logger.Infof("bmwFactory creates moto, model: '%s' ", model)
	return f.motoCreator(model, speed, f.logger)
}


// NewBMWFactory creates new car factory
func NewBMWFactory(carCreator carCreator, motoCreator motoCreator, logger *logrus.Logger) Factory {
	return &bmwFactory{
		carCreator: carCreator,
		motoCreator: motoCreator,
		logger:     logger,
	}
}
