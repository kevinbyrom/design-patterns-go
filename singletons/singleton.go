package singletons

import (
	"sync"
)

var mutex = &sync.Mutex{}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(singleton)
		}
		mutex.Unlock()
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
