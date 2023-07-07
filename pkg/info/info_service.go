package info

import (
	"go.uber.org/zap"
	"mcs-app-info/config/app"
)

type Servicer interface {
	GetAppInfo() (*Info, error)
}

type Service struct {
	logger     *zap.Logger
	commitHash string
}

func NewInfoService(logger *zap.Logger, commitHash string) *Service {
	return &Service{logger, commitHash}
}

func (service *Service) GetAppInfo() (*Info, error) {
	// this block make GetAppInfo() not testable.
	info := &Info{
		AppName:     app.Config.AppInfo.Name,
		Description: app.Config.AppInfo.Description,
		GitCommit:   service.commitHash,
		Version:     service.commitHash,
	}

	return info, nil
}
