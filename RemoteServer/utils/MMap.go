package bilicoin

import "sync"

type MMap struct {
	sync.RWMutex
	m map[string]string
	// m sessionID deviceID
}

func (s *MMap) Put(key string, value string) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (s *MMap) Get(key string) string {
	s.RLock()
	defer s.RUnlock()
	return s.m[key]
}

func NewMMap(capacity int) MMap {
	return MMap{
		m: make(map[string]string, capacity),
	}
}
