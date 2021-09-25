package pool

// Pool defines the interface for a goroutine pool
type Pool interface {
	// Go submits a function to the pool. The function represents a unit of work
	// that the pool will complete in the background
	Go(work func())
	// Wait blocks the calling thread until all jobs in the pool are completed
	Wait()
}

// WithPriority creates a new Pool instance with the given priority value. Priority is
// effectively just a uint value which determines the number of jobs to execute concurrently
// within the pool.
func WithPriority(priority Priority) Pool {
	return New(uint(priority))
}

// Default creates a new Pool instance with the default priorit.
func Default() Pool {
	return WithPriority(NORMAL)
}

// New creates a new Pool that can create up to `max` worker goroutines in parallel
func New(max uint) Pool {

	// Max cannot be zero. It must be at least one
	if max < 1 {
		max = 1
	}

	// Create a channel for available worker goroutines
	c := make(chan bool, max)
	for i := uint(0); i < max; i++ {
		c <- true
	}

	// Create and return the pool
	return &mPool{
		c: c,
	}

}
