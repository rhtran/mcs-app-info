package info

type Servicer interface {
	GetAppInfo() (*Info, error)
}

type Service struct {
	CommitHash string
	Info       Info
}

func NewInfoService() *Service {
	return &Service{}
}

func (service *Service) GetAppInfo() (*Info, error) {
	return &service.Info, nil
}
