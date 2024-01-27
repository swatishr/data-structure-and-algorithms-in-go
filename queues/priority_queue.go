package queues

import "errors"

// Implement priority queue with linked list
// Priority Queue is efficiently implemented through heaps but it will be covered in the heap section

// Using LL, push or pop: one operation will take O(n) time. So, depending upon your application
// if it's ready-heavy or write-heavy, you can try to optimize the corresponding operation.
// Below assumes a read-heavy application, so optimizing pop/dequeue to be O(1) and push/enqueue to be O(n)

type pNode struct {
	val      int
	priority int
	next     *pNode
}

type priorityQueue struct {
	front *pNode
	rear  *pNode
	len   int
}

func (q *priorityQueue) enqueue(x int, priority int) {
	curr := &pNode{
		val:      x,
		priority: priority,
		next:     nil,
	}

	if q.isEmpty() {
		q.front = curr
		q.rear = curr
		q.len += 1
		return
	}

	ptr := q.front
	var prev *pNode
	for ptr != nil {
		if ptr.priority < curr.priority {
			break
		}
		prev = ptr
		ptr = ptr.next
	}
	if prev == nil {
		curr.next = q.front
		q.front = curr
	} else {
		curr.next = prev.next
		prev.next = curr
	}

	if q.rear.next != nil {
		q.rear = q.rear.next
	}

	q.len += 1
}

func (q *priorityQueue) dequeue() int {
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

func (q *priorityQueue) getFront() int {
	if q.isEmpty() {
		panic(errors.New("queue is empty"))
	}
	return q.front.val
}

func (q *priorityQueue) getRear() int {
	if q.isEmpty() {
		panic(errors.New("queue is empty"))
	}
	return q.rear.val
}

func (q *priorityQueue) isEmpty() bool {
	return q.front == nil && q.rear == nil
}

func (q *priorityQueue) length() int {
	return q.len
}
