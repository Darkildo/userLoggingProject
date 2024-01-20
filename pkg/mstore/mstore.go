package mstore

import "errors"

type Mstore[T comparable] struct {
	store *SafetySlice[T]
}

func NewStore[T comparable](dataSlice *[]T) *Mstore[T] {
	return &Mstore[T]{store: NewSlice(dataSlice)}
}
func (s *Mstore[T]) Add(value ...T) (int, error) {

	s.store.mx.Lock()
	s.store.data = append(s.store.data, value...)
	s.store.mx.Unlock()
	return len(s.store.data) - 1, nil

}
func (s *Mstore[T]) Get(index int) (any, error) {
	if len(s.store.data) <= index {
		return nil, errors.New("invalid index")
	}
	s.store.mx.RLock()
	defer s.store.mx.RUnlock()
	return s.store.data[index], nil
}
func (s *Mstore[T]) Update(index int, value T) (int, error) {
	if len(s.store.data) <= index {
		return -1, errors.New("invalid index")
	}
	s.store.mx.Lock()

	s.store.data[index] = value

	s.store.mx.Unlock()

	return index, nil
}
func (s *Mstore[T]) Remove(value T) error {
	if len(s.store.data) == 0 {
		return errors.New("store is empty")
	}
	s.store.mx.Lock()
	defer s.store.mx.Unlock()
	findIndex := -1
	for i, datum := range s.store.data {
		if datum == value {
			findIndex = i
			break
		}
	}
	if len(s.store.data) == 0 {
		return errors.New("element not found")
	}
	s.store.data = append(s.store.data[:findIndex], s.store.data[findIndex+1:]...)

	return nil
}
func (s *Mstore[T]) RemoveByIndex(index int) error {
	if len(s.store.data) == 0 {
		return errors.New("store is empty")
	}
	s.store.mx.Lock()
	defer s.store.mx.Unlock()

	if len(s.store.data) <= index {
		return errors.New("element not found")
	}
	s.store.data = append(s.store.data[:index], s.store.data[index+1:]...)

	return nil
}
func (s *Mstore[T]) GetAll() []T {

	s.store.mx.RLock()
	defer s.store.mx.RUnlock()
	return s.store.data
}

func (s *Mstore[T]) Clear() {

	s.store.mx.Lock()
	defer s.store.mx.Unlock()
	s.store.data = make([]T, 0)

}
