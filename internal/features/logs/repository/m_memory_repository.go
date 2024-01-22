package repository

import (
	"errors"
	"github.com/lrita/cmap"
	LogEntry "userLoggingProject/internal/features/logs/entity"
	"userLoggingProject/pkg/mstore"
)

type MMemoryRepository struct {
	store *cmap.Map[string, *mstore.Mstore[LogEntry.LogEntry]]
}

func NewMemoryRepo() *MMemoryRepository {

	return &MMemoryRepository{store: &cmap.Map[string, *mstore.Mstore[LogEntry.LogEntry]]{}}
}

func (s *MMemoryRepository) SaveAll(userId string, entries []LogEntry.LogEntry) error {

	var lEntries *mstore.Mstore[LogEntry.LogEntry]
	var ok bool
	lEntries, ok = s.store.Load(userId)
	if !ok {
		var slice []LogEntry.LogEntry
		lEntries = mstore.NewStore[LogEntry.LogEntry](&slice)
	}

	_, err := lEntries.Add(entries...)
	if err != nil {
		return err
	}
	s.store.Store(userId, lEntries)
	return nil
}
func (s *MMemoryRepository) Save(userId string, entry *LogEntry.LogEntry) (int, error) {
	var entries *mstore.Mstore[LogEntry.LogEntry]
	var ok bool
	entries, ok = s.store.Load(userId)
	if !ok {
		var slice []LogEntry.LogEntry
		entries = mstore.NewStore[LogEntry.LogEntry](&slice)
	}

	i, err := entries.Add(*entry)
	if err != nil {
		return -1, err
	}
	s.store.Store(userId, entries)
	return i, nil
}
func (s *MMemoryRepository) LoadAll(userId string) ([]LogEntry.LogEntry, error) {

	entries, ok := s.store.Load(userId)

	if !ok {
		return make([]LogEntry.LogEntry, 0), nil
	}

	return entries.GetAll(), nil
}
func (s *MMemoryRepository) Load(userId string, count int) ([]LogEntry.LogEntry, error) {
	st, ok := s.store.Load(userId)

	if !ok {
		return make([]LogEntry.LogEntry, 0), nil
	}
	entries := st.GetAll()
	if len(entries) <= count {
		return entries, nil
	}
	return entries[:count], nil
}
func (s *MMemoryRepository) RemoveAll(userId string) error {
	st, ok := s.store.Load(userId)

	if !ok {
		return nil
	}
	st.Clear()

	s.store.Store(userId, st)

	s.store.Delete(userId)

	return nil
}
func (s *MMemoryRepository) Remove(userId string, logId int) error {
	st, ok := s.store.Load(userId)

	if !ok {
		return errors.New("element  not found")
	}

	err := st.RemoveByIndex(logId)
	if err != nil {
		return err
	}

	return nil
}
