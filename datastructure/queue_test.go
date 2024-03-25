package datastructure

import (
	"testing"

	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

func Test_Queue(t *testing.T) {
	type MODE int
	const (
		ENQUEUE = MODE(iota + 1)
		DEQUEUE
		IS_FULL
		IS_EMPTY
		// FRONT
		// ITEMS
	)

	type (
		want struct {
			length  int
			front   int
			isEmpty bool
			isFull  bool
		}
		args struct {
			item int
		}
		test struct {
			name       string
			mode       MODE
			queue      QueueInterface[int]
			args       args
			want       want
			beforeFunc func(queue QueueInterface[int]) QueueInterface[int]
			afterFunc  func(queue QueueInterface[int]) QueueInterface[int]
		}
	)

	tests := []test{
		{
			name:  "enqueue",
			mode:  ENQUEUE,
			queue: NewQueue[int](),
			args:  args{item: 1},
			want:  want{length: 1, front: 1},
			afterFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				t.Logf("items: %+v", queue.Items())
				return queue
			},
		},
		{
			name:  "dequeue",
			mode:  DEQUEUE,
			queue: NewQueue[int](),
			want:  want{length: 1, front: 2},
			beforeFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				queue.Enqueue(1)
				queue.Enqueue(2)
				return queue
			},
			afterFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				t.Logf("items: %+v", queue.Items())
				return queue
			},
		},
		{
			name:  "is empty",
			mode:  IS_EMPTY,
			queue: NewQueue[int](),
			want:  want{isEmpty: true},
			afterFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				t.Logf("items: %+v", queue.Items())
				return queue
			},
		},
		{
			name:  "is empty",
			mode:  IS_EMPTY,
			queue: NewQueue[int](),
			want:  want{isEmpty: false},
			beforeFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				queue.Enqueue(1)
				return queue
			},
			afterFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				t.Logf("items: %+v", queue.Items())
				return queue
			},
		},
		{
			name:  "is full",
			mode:  IS_FULL,
			queue: NewQueue[int]().WithMax(1),
			want:  want{isFull: true},
			beforeFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				queue.Enqueue(1)
				queue.Enqueue(2)
				return queue
			},
			afterFunc: func(queue QueueInterface[int]) QueueInterface[int] {
				t.Logf("items: %+v", queue.Items())
				return queue
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeFunc != nil {
				tt.queue = tt.beforeFunc(tt.queue)
			}

			switch tt.mode {
			case ENQUEUE:
				if err := tt.queue.Enqueue(tt.args.item); err != nil {
					t.Fatalf("error while enqueue, %v", err)
				}

				if tt.want.length != len(tt.queue.Items()) {
					t.Fatalf("want length item is '%v' but got '%v'", tt.want.length, len(tt.queue.Items()))
				}

				front, err := tt.queue.Front()
				if err != nil && errors.GetCode(err) != codes.CodeQueueEmpty {
					t.Fatalf("error while get front queue, %v", err)
				}

				if tt.want.front != front {
					t.Fatalf("want front item is '%v' but got '%v'", tt.want.front, front)
				}
			case DEQUEUE:
				if _, err := tt.queue.Dequeue(); err != nil {
					t.Fatalf("error while dequeue, %v", err)
				}

				if tt.want.length != len(tt.queue.Items()) {
					t.Fatalf("want length item is '%v' but got '%v'", tt.want.length, len(tt.queue.Items()))
				}

				front, err := tt.queue.Front()
				if err != nil && errors.GetCode(err) != codes.CodeQueueEmpty {
					t.Fatalf("error while get front queue, %v", err)
				}

				if tt.want.front != front {
					t.Fatalf("want front item is '%v' but got '%v'", tt.want.front, front)
				}
			case IS_FULL:
				if full := tt.queue.IsFull(); full != tt.want.isFull {
					t.Fatalf("want is full is '%v' but got '%v'", tt.want.isFull, full)
				}
			case IS_EMPTY:
				if empty := tt.queue.IsEmpty(); empty != tt.want.isEmpty {
					t.Fatalf("want is empty is '%v' but got '%v'", tt.want.isEmpty, empty)
				}
			default:
				t.Fatalf("mode not supported!")
			}

			if tt.afterFunc != nil {
				tt.queue = tt.afterFunc(tt.queue)
			}
		})
	}
}
