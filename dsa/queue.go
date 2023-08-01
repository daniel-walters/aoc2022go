package dsa

import "errors"

type Queue[T any] []T

var ErrQueueEmpty = errors.New("Queue is empty")

func NewQueue[T any]() Queue[T] {
	return []T{}
}

func NewQueueFrom[T any](init []T) Queue[T] {
	return init
}

func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), ErrQueueEmpty
	}

	dequeued := (*q)[0]

	*q = (*q)[1:]

	return dequeued, nil
}

func (q Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), ErrQueueEmpty
	}

	return q[0], nil
}
