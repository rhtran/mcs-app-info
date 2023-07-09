package info

import (
	"github.com/ryyt-cg/mcs-app-info/config/app"
	"go.uber.org/zap"
)

type Servicer interface {
	GetAppInfo() (*Info, error)
}

type Service struct {
	logger     *zap.Logger
	commitHash string
	info Info
}

func NewInfoService(logger *zap.Logger, commitHash string, info Info) *Service {
	return &Service{logger, commitHash, info}
}

func (service *Service) GetAppInfo() (*Info, error) {
	return &service.info, nil
}
