package datastructure

import (
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/convert"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

type Queue[ITEM any] struct {
	max   *int
	items []ITEM
}

type QueueInterface[ITEM any] interface {
	WithMax(max int) QueueInterface[ITEM]
	Enqueue(item ITEM) error
	Dequeue() (ITEM, error)
	IsEmpty() bool
	IsFull() bool
	Front() (ITEM, error)
	Items() []ITEM
}

func NewQueue[ITEM any]() QueueInterface[ITEM] {
	return &Queue[ITEM]{
		items: []ITEM{},
	}
}

// Items implements QueueInterface.
func (q *Queue[ITEM]) Items() []ITEM {
	return q.items
}

// Dequeue implements QueueInterface.
func (q *Queue[ITEM]) Dequeue() (ITEM, error) {
	if len(q.items) <= 0 {
		return convert.ToSafeValue[ITEM](new(ITEM)), errors.NewWithCode(codes.CodeQueueEmpty, "queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Enqueue implements QueueInterface.
func (q *Queue[ITEM]) Enqueue(item ITEM) error {
	if q.max != nil {
		max := convert.ToSafeValue[int](q.max)
		items := len(q.items)
		if items >= max {
			return errors.NewWithCode(codes.CodeQueueFull, "maximum queue is %v and current item length is %v", max, items)
		}
	}

	q.items = append(q.items, item)
	return nil
}

// Front implements QueueInterface.
func (q *Queue[ITEM]) Front() (ITEM, error) {
	if len(q.items) <= 0 {
		return convert.ToSafeValue[ITEM](new(ITEM)), errors.NewWithCode(codes.CodeQueueEmpty, "queue is empty")
	}

	front := q.items[0]
	return front, nil
}

// IsEmpty implements QueueInterface.
func (q *Queue[ITEM]) IsEmpty() bool {
	return len(q.items) == 0
}

// IsFull implements QueueInterface.
func (q *Queue[ITEM]) IsFull() bool {
	if q.max != nil {
		if len(q.items) >= convert.ToSafeValue[int](q.max) {
			return true
		}
	}
	return false
}

// WithMax implements QueueInterface.
func (q *Queue[ITEM]) WithMax(max int) QueueInterface[ITEM] {
	q.max = &max
	return q
}
