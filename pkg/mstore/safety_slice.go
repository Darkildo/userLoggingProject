package mstore

import (
	"sync"
)

type SafetySlice[T comparable] struct {
	data []T
	mx   sync.RWMutex
}

func NewSlice[T comparable](data *[]T) *SafetySlice[T] {
	return &SafetySlice[T]{data: *data, mx: sync.RWMutex{}}
}
