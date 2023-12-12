package maps

import (
	"sync"
)

func WriteMapNoSync(m map[string]string, k, v string) {
	m[k] = v
}

func WriteMapMutex(m map[string]string, k, v string, mu *sync.Mutex) {
	mu.Lock()
	m[k] = v
	mu.Unlock()
}

func WriteSyncMap(m *sync.Map, k, v string) {
	m.Store(k, v)
}
