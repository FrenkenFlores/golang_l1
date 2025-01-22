package l1_18

import "sync"

type CountedStruct struct {
	i  int
	mu sync.Mutex
}

func (s *CountedStruct) Inc() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.i++
}

func (s *CountedStruct) Get() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.i
}
