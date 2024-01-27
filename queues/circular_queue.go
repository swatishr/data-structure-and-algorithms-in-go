package queues

import "errors"

// This covers an implementation of a circular queue using arrays.
// You can use linked list too with a given capacity.
// We are using arrays here to cover all the different variants of queue implementation.
// There is another variant of circular queue where when the queue is full, you increase the capacity of the array then.
// For eg. double the capacity.

type circularQueue struct {
	arr      []int
	len      int
	capacity int
	front    int
	rear     int
}

func NewCircularQueue(capacity int) *circularQueue {
	cq := &circularQueue{
		capacity: capacity,
		front:    -1,
		rear:     -1,
	}
	cq.arr = make([]int, capacity)
	return cq
}

func (q *circularQueue) enqueue(x int) {
	if q.isFull() {
		panic(errors.New("queue is full"))
	}

	q.rear = (q.rear + 1) % q.capacity
	q.arr[q.rear] = x
	if q.front == -1 {
		q.front = q.rear
	}
	q.len += 1
}

func (q *circularQueue) dequeue() int {
	if q.isEmpty() {
		panic(errors.New("no data to dequeue"))
	}
	dqVal := q.arr[q.front]

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % q.capacity
	}

	q.len -= 1

	return dqVal
}

func (q *circularQueue) getFront() int {
	if q.isEmpty() {
		panic(errors.New("queue is empty"))
	}
	return q.arr[q.front]
}

func (q *circularQueue) getRear() int {
	if q.isEmpty() {
		panic(errors.New("queue is empty"))
	}
	return q.arr[q.rear]
}

func (q *circularQueue) isEmpty() bool {
	return q.front == -1 && q.rear == -1
}

func (q *circularQueue) isFull() bool {
	return q.front == (q.rear+1)%q.capacity
}

func (q *circularQueue) size() int {
	return q.len
}
