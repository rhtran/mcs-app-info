package health

import (
	"go.uber.org/zap"
)

type CheckServicer interface {
	Check() (*Check, error)
}

type CheckService struct {
	logger *zap.Logger
}

func NewHealthCheckService(logger *zap.Logger) *CheckService {
	return &CheckService{logger}
}

func (service *CheckService) Check() (*Check, error) {
	healthCheck := &Check{
		Status: "UP",
	}

	return healthCheck, nil
}
