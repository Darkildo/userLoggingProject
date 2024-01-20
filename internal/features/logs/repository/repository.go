package repository

import LogEntry "userLoggingProject/internal/features/logs/entity"

type LogsRepository interface {
	SaveAll(userId string, entries []LogEntry.LogEntry) error
	Save(userId string, entry *LogEntry.LogEntry) (int, error)
	LoadAll(userId string) ([]LogEntry.LogEntry, error)
	Load(userId string, count int) ([]LogEntry.LogEntry, error)
	RemoveAll(userId string) error
	Remove(userId string, logId int) error
}
