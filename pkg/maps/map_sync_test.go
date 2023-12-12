package maps_test

import (
	"sync"
	"testing"

	"github.com/nyanshak/golang-concurrency-demo/pkg/maps"
)

func TestWriteMapNoSync(t *testing.T) {
	m := map[string]string{}

	t.Run("case 1", func(t *testing.T) {
		t.Parallel()
		maps.WriteMapNoSync(m, "key", "value1")
	})

	t.Run("case 2", func(t *testing.T) {
		t.Parallel()
		maps.WriteMapNoSync(m, "key", "value2")
	})
}

func TestWriteMapMutex(t *testing.T) {
	m := map[string]string{}
	mutex := &sync.Mutex{}

	t.Run("case 1", func(t *testing.T) {
		t.Parallel()
		maps.WriteMapMutex(m, "key", "value1", mutex)
	})

	t.Run("case 2", func(t *testing.T) {
		t.Parallel()
		maps.WriteMapMutex(m, "key", "value2", mutex)
	})
}

func TestWriteSyncMap(t *testing.T) {
	var m sync.Map

	t.Run("case 1", func(t *testing.T) {
		t.Parallel()
		maps.WriteSyncMap(&m, "key", "value1")
	})

	t.Run("case 2", func(t *testing.T) {
		t.Parallel()
		maps.WriteSyncMap(&m, "key", "value2")
	})
}
