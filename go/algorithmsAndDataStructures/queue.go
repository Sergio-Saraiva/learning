package main

type Node[T any] struct {
	value T
	next *Node[T]
}

type Queue[T any] struct {
	length int32
	head *Node[T]
	tail *Node[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		length: 0,
		head: nil,
		tail: nil,
	}
}

func (q *Queue[T]) enqueue(item T) {
	newNode := Node[T] {
		value: item,
		};
		
	if(q.length == 0) {
		q.length++;
		q.tail = &newNode;
		q.head = q.tail;
		return
	}
	q.length++;
	q.tail.next = &newNode;
	q.tail = &newNode;
	return
}

func (q *Queue[T]) deque() T {
	if (q.length == 0) {
		var zero T
		return zero;
	}

	q.length--;

	head := q.head;
	q.head = q.head.next;
	head.next = nil;
	return head.value;
}

func (q *Queue[T]) Peek() T {
	if(q.length == 0) {
		var zero T
		return zero
	}
	
	return q.head.value;
}