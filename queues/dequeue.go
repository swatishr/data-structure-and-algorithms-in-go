package queues

import (
	"errors"
)

// Double-ended queue can be implemented either using doubly-linked list or circular list
// here, we are using doubly-linked list

type dLNode struct {
	val  int
	next *dLNode
	prev *dLNode
}

type dequeue struct {
	front *dLNode
	rear  *dLNode
	len   int
}

func (q *dequeue) pushFront(x int) {
	curr := &dLNode{
		val:  x,
		next: nil,
		prev: nil,
	}
	if q.isEmpty() {
		q.rear = curr
	} else {
		q.front.prev = curr
		curr.next = q.front
	}
	q.front = curr
	q.len += 1
}

func (q *dequeue) pushBack(x int) {
	curr := &dLNode{
		val:  x,
		next: nil,
		prev: nil,
	}
	if q.isEmpty() {
		q.front = curr
	} else {
		q.rear.next = curr
		curr.prev = q.rear
	}
	q.rear = curr
	q.len += 1
}

func (q *dequeue) popFront() int {
	if q.isEmpty() {
		panic(errors.New("no data to pop"))
	}
	dqVal := q.front.val
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	} else {
		q.front.prev = nil
	}

	q.len -= 1

	return dqVal
}

func (q *dequeue) popBack() int {
	if q.isEmpty() {
		panic(errors.New("no data to pop"))
	}
	dqVal := q.rear.val
	q.rear = q.rear.prev

	if q.rear == nil {
		q.front = nil
	} else {
		q.rear.next = nil
	}

	q.len -= 1

	return dqVal
}

func (q *dequeue) getFront() int {
	if q.isEmpty() {
		panic(errors.New("dequeue is empty"))
	}
	return q.front.val
}

func (q *dequeue) getRear() int {
	if q.isEmpty() {
		panic(errors.New("dequeue is empty"))
	}
	return q.rear.val
}

func (q *dequeue) isEmpty() bool {
	return q.front == nil && q.rear == nil
}

func (q *dequeue) size() int {
	return q.len
}
