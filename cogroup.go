package cogroup

import (
	"sync"
)

// CoGroup wraps WaitGroup and adds a returnable counter.
type CoGroup struct {
	sync.RWMutex
	sync.WaitGroup
	c int
}

// Add increases the counter.
func (co *CoGroup) Add(n int) {
	co.Lock()
	defer co.Unlock()
	co.c += n
	co.WaitGroup.Add(n)
}

// Count returns the current count.
func (co *CoGroup) Count() int {
	co.RLock()
	defer co.RUnlock()
	return co.c
}

// Done decrements the CoGroup counter by one.
func (co *CoGroup) Done() {
	co.Lock()
	defer co.Unlock()
	co.c--
	co.WaitGroup.Done()
}
