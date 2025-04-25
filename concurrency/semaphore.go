package concurrency

import "sync"

type (
	SemaphoreOpts struct {
		Limit int
	}

	semaphore struct {
		c chan struct{}
	}
)

// Default set default value for semaphore options if not set
func (so *SemaphoreOpts) Default() {
	if so.Limit <= 0 {
		so.Limit = 1
	}
}

// NewSemaphoreLocker return sync.Locker with custom limit of concurrency
// Semaphore binary implementation of sync.Locker with channel
func NewSemaphoreLocker(opts SemaphoreOpts) sync.Locker {
	opts.Default()

	l := &semaphore{
		c: make(chan struct{}, opts.Limit),
	}

	for i := 0; i < opts.Limit; i++ {
		l.c <- struct{}{}
	}

	return l
}

// Lock acquire a lock
func (l *semaphore) Lock() {
	<-l.c
}

// Unlock release a lock
func (l *semaphore) Unlock() {
	l.c <- struct{}{}
}
