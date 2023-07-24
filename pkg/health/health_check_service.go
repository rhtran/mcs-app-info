package health

type CheckServicer interface {
	Check() (*Check, error)
}

type CheckService struct{}

func NewHealthCheckService() *CheckService {
	return &CheckService{}
}

func (service *CheckService) Check() (*Check, error) {
	healthCheck := &Check{
		Status: "UP",
	}

	return healthCheck, nil
}
