package utils

// Generic Queue using slice
type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

func (q *Queue[T]) Push(v T) {
	q.items = append(q.items, v)
}

func (q *Queue[T]) Pop() T {
	v := q.items[0]
	q.items = q.items[1:]
	return v
}

func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}
