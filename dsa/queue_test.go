package dsa_test

import (
	"aoc/dsa"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := dsa.NewQueue[int]()

	if !queue.IsEmpty() {
		t.Error()
	}

	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)

	if v, _ := queue.Peek(); v != 4 {
		t.Error()
	}

	if v, _ := queue.Dequeue(); v != 4 {
		t.Error()
	}

	if v, _ := queue.Peek(); v != 5 {
		t.Error()
	}

	if queue.IsEmpty() {
		t.Error()
	}
}
