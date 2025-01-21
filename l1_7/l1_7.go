package l1_7

import "sync"

type MutexMap struct {
	mut  sync.Mutex
	Data map[string]any
}

func (m *MutexMap) Read(key string) (any, bool) {
	m.mut.Lock()
	defer m.mut.Unlock()
	value, ok := m.Data[key]
	return value, ok
}

func (m *MutexMap) Write(key string, value any) {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.Data[key] = value

}
