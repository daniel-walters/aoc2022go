package dsa

import "errors"

type Stack[T any] []T

var ErrStackEmpty = errors.New("Stack is empty")

func NewStack[T any]() Stack[T] {
	return []T{}
}

func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), ErrStackEmpty
	}

	popped := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return popped, nil
}

func (s Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), ErrStackEmpty
	}

	return s[len(s)-1], nil
}
