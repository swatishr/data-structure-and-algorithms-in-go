package queues

import "errors"

type node struct {
	val  int
	next *node
}

type queue struct {
	front *node
	rear  *node
	len   int
}

func (q *queue) enqueue(x int) {
	curr := &node{
		val:  x,
		next: nil,
	}
	if q.isEmpty() {
		q.front = curr
	} else {
		q.rear.next = curr
	}
	q.rear = curr
	q.len += 1
}

func (q *queue) dequeue() int {
	if q.isEmpty() {
		panic(errors.New("no data to dequeue"))
	}
	dqVal := q.front.val
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.len -= 1

	return dqVal
}

func (q *queue) getFront() int {
	if q.isEmpty() {
		panic(errors.New("queue is empty"))
	}
	return q.front.val
}

func (q *queue) getRear() int {
	if q.isEmpty() {
		panic(errors.New("queue is empty"))
	}
	return q.rear.val
}

func (q *queue) isEmpty() bool {
	return q.front == nil && q.rear == nil
}

func (q *queue) length() int {
	return q.len
}
