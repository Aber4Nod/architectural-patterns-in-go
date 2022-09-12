package moto

import "github.com/sirupsen/logrus"

type moto struct {
	model  string
	speed  int
	logger *logrus.Logger
}

// Forward commands moto to move forward
func (m *moto) Forward() {
	m.logger.Infof("%s moves froward with speed %d", m.model, m.speed)
}

// Back commands moto to move back
func (m *moto) Back() {
	m.logger.Infof("%s moves back with speed %d", m.model, m.speed)
}

// SomeAction2 commands moto to perform some action2
func (m *moto) SomeAction2() {
	m.logger.Infof("%s performs some action2 with speed %d", m.model, m.speed)
}

// GetModel returns product model
func (m *moto) GetModel() string {
	return m.model
}

// GetSpeed returns product speed
func (m *moto) GetSpeed() int {
	return m.speed
}

// NewMoto creates new moto
func NewMoto(model string, speed int, logger *logrus.Logger) Moto {
	return &moto{
		model:  model,
		speed:  speed,
		logger: logger,
	}
}
