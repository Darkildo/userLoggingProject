package service_provider

import "userLoggingProject/internal/features/logs/repository"

type ServiceProvider struct {
	Logs *repository.LogsRepository
}

func NewServiceProvider(logsRepo *repository.LogsRepository) *ServiceProvider {
	return &ServiceProvider{Logs: logsRepo}
}
