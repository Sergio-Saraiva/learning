package main

import (
	"math"
)

type StackNode[T any] struct {
	value T
	prev *StackNode[T]
}

type Stack[T any] struct {
	length int32;
	head *StackNode[T]

}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		length: 0,
		head: nil,
	}
}

func (s *Stack[T]) Push(item T) {
	newStackNode := StackNode[T] {
		value: item,
	}
	s.length++;
	if(s.head == nil) {
		s.head = &newStackNode;
		return
	}

	newStackNode.prev = s.head;
	s.head = &newStackNode;
}

func (s *Stack[T]) Pop() T {
	if(s.length == 0){
		s.head = nil
		var zero T
		return zero
	}
	s.length = int32(math.Max(0, float64(s.length) - 1))
	head := s.head
	s.head = head.prev
	return head.value
}

func (s *Stack[T]) Peek() T {
	if(s.length == 0){
		var zero T
		return zero
	}
	return s.head.value
}