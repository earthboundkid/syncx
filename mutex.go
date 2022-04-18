package syncx

import "sync"

type Mutex[T any] struct {
	mu    sync.RWMutex
	value T
}

func NewMutex[T any](initial T) *Mutex[T] {
	var m Mutex[T]
	m.value = initial
	return &m
}

func (m *Mutex[T]) ReadLock(f func(value T)) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	f(m.value)
}

func (m *Mutex[T]) Lock(f func(value *T)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value := m.value
	f(&value)
	m.value = value
}

func (m *Mutex[T]) Read() T {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.value
}

func (m *Mutex[T]) Write(value T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}