package villa

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestOnce(t *testing.T) {
	counter := int32(0)
	o := Once{F: func() {
		atomic.AddInt32(&counter, 1)
	}}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			o.Do()
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, "counter", counter, int32(1))
}

func TestAtomicBox(t *testing.T) {
	var b AtomicBox

	b.Set("hello")
	assert.Equal(t, "b.Get", b.Get().(string), "hello")
}
