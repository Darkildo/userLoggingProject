package service_provider

import (
	"userLoggingProject/internal/features/logs/service"
)

type ServiceProvider struct {
	Logs *service.LogService
}

func NewServiceProvider(logsRepo *service.LogService) *ServiceProvider {
	return &ServiceProvider{Logs: logsRepo}
}
