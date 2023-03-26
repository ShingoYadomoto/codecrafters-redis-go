package store

import (
	"sync"
	"time"
)

type storeWithExpiry struct {
	mu       sync.Mutex
	valStore map[string]string
	expiry   map[string]time.Time
}

func (s *storeWithExpiry) Store(key, value string, expiry time.Duration) {
	s.mu.Lock()
	s.valStore[key] = value
	if expiry != 0 {
		s.expiry[key] = time.Now().Add(expiry)
	}
	s.mu.Unlock()
}

func (s *storeWithExpiry) Load(key string) (value string, ok bool) {
	s.mu.Lock()

	val, ok := s.valStore[key]
	if !ok {
		return "", false
	}

	if expireTime, ok := s.expiry[key]; ok {
		if expireTime.Before(time.Now()) {
			return "", false
		}
	}
	s.mu.Unlock()

	return val, true
}

var store = &storeWithExpiry{
	valStore: map[string]string{},
	expiry:   map[string]time.Time{},
}

func GetStore() *storeWithExpiry {
	return store
}
