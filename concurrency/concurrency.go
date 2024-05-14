package concurrency

import (
	"context"
	"sync"
)

type Interface interface {

	// Use your custom sync.WaitGroup
	WithWg(wg *sync.WaitGroup) Interface

	// Use your custom sync.Mutex
	WithLocker(lc sync.Locker) Interface

	// Set maximum worker here. Default is 1
	WithMaxWorker(maxWorker int64) Interface

	// Run the list functions with goroutine
	Do(ctx context.Context) error

	// Added function that will be run async at goroutine. This method already call c.Done() after process is complete
	AddFunc(fn func(ctx context.Context, c Interface))

	// Lock block of code. This like mt.Lock()
	Lock()

	// Unlock block of code. This like mt.Unlock()
	Unlock()

	// To flag wait group if proccess is done. This like wg.Done()
	Done()

	// Added errors if have error. This method already implement c.Lock() and c.Unlock(). The error will be returned at Do() method if exists. Recommended to use errors.NewWithCode()
	AddError(errs ...error)

	// To clear list functions
	ClearFunc()
}

type concurrency struct {
	wg        *sync.WaitGroup
	lc        sync.Locker
	maxWorker int64
	listErr   []error
	listFunc  []func(ctx context.Context, c Interface)
}

// Do concurrency proccess with custom maximum worker
func NewConcurrency() Interface {
	result := concurrency{
		wg:        &sync.WaitGroup{},
		lc:        &sync.Mutex{},
		maxWorker: 1,
	}
	return &result
}

func (c *concurrency) WithWg(wg *sync.WaitGroup) Interface {
	c.wg = wg
	return c
}

func (c *concurrency) WithLocker(lc sync.Locker) Interface {
	c.lc = lc
	return c
}

func (c *concurrency) WithMaxWorker(maxWorker int64) Interface {
	c.maxWorker = maxWorker
	return c
}

func (c *concurrency) Do(ctx context.Context) error {
	worker := 0
	lenDo := len(c.listFunc)
	for i, fn := range c.listFunc {
		worker += 1
		c.wg.Add(1)
		go fn(ctx, c)

		if worker >= int(c.maxWorker) || i == (lenDo-1) {
			worker = 0
			c.wg.Wait()
			for _, err := range c.listErr {
				return err
			}
		}
	}

	c.ClearFunc()
	return nil
}

func (c *concurrency) AddFunc(fn func(ctx context.Context, c Interface)) {
	c.listFunc = append(c.listFunc, func(ctx context.Context, c Interface) {
		defer c.Done()
		fn(ctx, c)
	})
}

func (c *concurrency) Lock() {
	c.lc.Lock()
}

func (c *concurrency) Unlock() {
	c.lc.Unlock()
}

func (c *concurrency) Done() {
	c.wg.Done()
}

func (c *concurrency) AddError(errs ...error) {
	c.Lock()
	c.listErr = append(c.listErr, errs...)
	c.Unlock()
}

func (c *concurrency) ClearFunc() {
	c.listFunc = nil
}
