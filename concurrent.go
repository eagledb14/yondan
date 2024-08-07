package main

import (
	"sync"
	"errors"
)

type ConcurrentMap[V any] struct {
	data map[string]V
	mutex sync.RWMutex
}

func NewConcurrentMap[V any]() *ConcurrentMap[V] {
	newConcurrentMap := ConcurrentMap[V] {
		data: make(map[string]V),
		mutex: sync.RWMutex{},
	}

	return &newConcurrentMap
}

func (c *ConcurrentMap[V]) Read(target string) (V, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	run, ok := c.data[target]
	if !ok {
		var zero V
		return zero, errors.New("Target Not Found")
	}
	return run, nil
}

func (c *ConcurrentMap[V]) Write(target string, run V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[target] = run
}

func (c *ConcurrentMap[V]) MassWrite(newMap *map[string]V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for target, value := range *newMap {
		c.data[target] = value
	}
}

func (c *ConcurrentMap[V]) Delete(target string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, target)
}
