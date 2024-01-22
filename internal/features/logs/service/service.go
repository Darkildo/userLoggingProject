package service

import (
	LogEntry "userLoggingProject/internal/features/logs/entity"
	"userLoggingProject/internal/features/logs/repository"
)

type LogService struct {
	repo repository.LogsRepository
}

func newService(repo repository.LogsRepository) *LogService {
	return &LogService{repo: repo}
}

func (s *LogService) AddLog(userId string, log *LogEntry.LogEntry) (int, error) {
	return s.repo.Save(userId, log)
}

func (s *LogService) RemoveLog(userId string, logId int) error {
	return s.repo.Remove(userId, logId)
}

func (s *LogService) ClearLogs(userId string) error {
	return s.ClearLogs(userId)
}

func (s *LogService) GetAll(userId string) ([]LogEntry.LogEntry, error) {
	return s.GetAll(userId)
}
func (s *LogService) GetById(userId string, logId int) ([]LogEntry.LogEntry, error) {
	return s.GetById(userId, logId)
}
