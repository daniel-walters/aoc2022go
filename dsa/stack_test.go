package dsa_test

import (
	"aoc/dsa"
	"testing"
)

func TestStack(t *testing.T) {
	stack := dsa.NewStack[int]()

	if !stack.IsEmpty() {
		t.Error()
	}

	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	if v, _ := stack.Peek(); v != 6 {
		t.Error()
	}

	if v, _ := stack.Pop(); v != 6 {
		t.Error()
	}

	if v, _ := stack.Peek(); v != 5 {
		t.Error()
	}

	if stack.IsEmpty() {
		t.Error()
	}
}
