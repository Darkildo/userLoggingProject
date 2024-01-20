package repository

import (
	"errors"
	"github.com/lrita/cmap"
	LogEntry "userLoggingProject/internal/features/logs/entity"
)

type MMemoryRepository struct {
	mstore *cmap.Map[string, []LogEntry.LogEntry]
}

func (s *MMemoryRepository) SaveAll(userId string, entries []LogEntry.LogEntry) error {
	var lEntries []LogEntry.LogEntry
	lEntries, _ = s.mstore.Load(userId)
	lEntries = append(entries, entries...)

	return nil
}
func (s *MMemoryRepository) Save(userId string, entry *LogEntry.LogEntry) (int, error) {
	var entries []LogEntry.LogEntry
	entries, _ = s.mstore.Load(userId)

	entries = append(entries, *entry)

	return len(entries) - 1, nil
}
func (s *MMemoryRepository) LoadAll(userId string) ([]LogEntry.LogEntry, error) {

	entries, ok := s.mstore.Load(userId)

	if !ok {
		return make([]LogEntry.LogEntry, 0), nil
	}

	return entries, nil
}
func (s *MMemoryRepository) Load(userId string, count int) ([]LogEntry.LogEntry, error) {
	entries, ok := s.mstore.Load(userId)

	if !ok {
		return make([]LogEntry.LogEntry, 0), nil
	}
	if len(entries) <= count {
		return entries, nil
	}
	return entries[:count], nil
}
func (s *MMemoryRepository) RemoveAll(userId string) error {
	entries, ok := s.mstore.Load(userId)

	if !ok {
		return nil
	}
	entries = nil
	s.mstore.Delete(userId)

	return nil
}
func (s *MMemoryRepository) Remove(userId string, logId int) error {
	entries, ok := s.mstore.Load(userId)

	if !ok {
		return errors.New("element  not found")
	}

	if len(entries) <= logId {
		return errors.New("element not found")
	}
	entries = append(entries[:logId], entries[logId+1:]...)
	return nil
}
