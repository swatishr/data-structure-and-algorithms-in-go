package queues

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCQEnqueue(t *testing.T) {

	q := NewCircularQueue(5)

	require.NotPanics(t, func() { q.enqueue(4) })
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.enqueue(5)
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 5, q.getRear())

	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	require.PanicsWithError(t, "queue is full", func() { q.enqueue(9) })
}

func TestCQDequeue(t *testing.T) {

	q := NewCircularQueue(5)

	require.PanicsWithError(t, "no data to dequeue", func() { q.dequeue() })

	q.enqueue(5)
	q.enqueue(4)

	require.Equal(t, 5, q.dequeue())
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(6)
	q.enqueue(7)
	require.PanicsWithError(t, "queue is full", func() { q.enqueue(9) })

	require.Equal(t, 4, q.dequeue())
	require.Equal(t, 3, q.getFront())
	require.Equal(t, 7, q.getRear())

	q.enqueue(9)
	require.Equal(t, 9, q.getRear())

	require.Equal(t, 3, q.dequeue())
	require.Equal(t, 2, q.dequeue())
	require.Equal(t, 6, q.dequeue())
	require.Equal(t, 7, q.dequeue())
	require.Equal(t, 9, q.dequeue())
	require.True(t, q.isEmpty())
	require.PanicsWithError(t, "no data to dequeue", func() { q.dequeue() })

}

func TestCQGetFront(t *testing.T) {
	q := NewCircularQueue(3)

	require.PanicsWithError(t, "queue is empty", func() { q.getFront() })

	q.enqueue(4)
	require.Equal(t, 4, q.getFront())

	q.enqueue(5)
	require.Equal(t, 4, q.getFront())

	q.dequeue()
	require.Equal(t, 5, q.getFront())
}

func TestCQGetRear(t *testing.T) {
	q := NewCircularQueue(3)

	require.PanicsWithError(t, "queue is empty", func() { q.getRear() })

	q.enqueue(4)
	require.Equal(t, 4, q.getRear())

	q.enqueue(5)
	require.Equal(t, 5, q.getRear())

	q.dequeue()
	require.Equal(t, 5, q.getRear())
}

func TestCQIsEmpty(t *testing.T) {
	q := NewCircularQueue(3)

	require.True(t, q.isEmpty())

	q.enqueue(4)
	q.enqueue(5)
	require.False(t, q.isEmpty())

	q.dequeue()
	q.dequeue()
	require.True(t, q.isEmpty())

}

func TestCQIsFull(t *testing.T) {
	q := NewCircularQueue(3)

	require.False(t, q.isFull())

	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	require.True(t, q.isFull())

	q.dequeue()
	require.False(t, q.isFull())

}

func TestCQSize(t *testing.T) {
	q := NewCircularQueue(3)

	require.Equal(t, 0, q.size())

	q.enqueue(4)
	q.enqueue(5)
	require.Equal(t, 2, q.size())

	q.enqueue(6)
	require.Equal(t, 3, q.size())

	q.dequeue()
	require.Equal(t, 2, q.size())

	q.enqueue(7)
	require.Equal(t, 3, q.size())

	q.dequeue()
	q.dequeue()
	q.dequeue()
	require.Equal(t, 0, q.size())
}
