package pool

import "sync"

// mPool is our internal implementation of the Pool interface.
type mPool struct {
	wg sync.WaitGroup
	c  chan bool
}

// Go submits a function to the pool. The function represents a unit of work
// that the pool will complete in the background
func (p *mPool) Go(work func()) {

	// Add one to the wait group
	p.wg.Add(1)

	// Avoid blocking the calling goroutine by spawning a new one to wait for
	// an available worker slow to open up
	go func() {

		// Wait for a slot to open up
		<-p.c

		// Defer a function to make the slot available
		defer func() {
			p.c <- true
			p.wg.Done()
		}()

		// Perform the work
		work()

	}()
}

// Wait blocks the calling thread until all jobs in the pool are completed
func (p *mPool) Wait() {

	// Wait for the wait group
	p.wg.Wait()

}
