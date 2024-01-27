package queues

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPushFront(t *testing.T) {

	q := dequeue{}

	require.NotPanics(t, func() { q.pushFront(4) })
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.pushFront(5)
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.pushFront(6)
	require.Equal(t, 6, q.getFront())
	require.Equal(t, 4, q.getRear())
}

func TestPushBack(t *testing.T) {

	q := dequeue{}

	require.NotPanics(t, func() { q.pushBack(4) })
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.pushBack(5)
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 5, q.getRear())

	q.pushBack(6)
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 6, q.getRear())
}

func TestPopFront(t *testing.T) {

	q := dequeue{}

	require.PanicsWithError(t, "no data to pop", func() { q.popFront() })

	q.pushBack(5)
	q.pushFront(4)

	require.Equal(t, 4, q.popFront())
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 5, q.getRear())

	q.pushFront(3)
	q.pushBack(2)

	require.Equal(t, 3, q.popFront())
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 2, q.getRear())

	require.Equal(t, 5, q.popFront())
	require.Equal(t, 2, q.getFront())
	require.Equal(t, 2, q.getRear())

	require.Equal(t, 2, q.popFront())

	require.True(t, q.isEmpty())
}

func TestPopBack(t *testing.T) {

	q := dequeue{}

	require.PanicsWithError(t, "no data to pop", func() { q.popFront() })

	q.pushBack(5)
	q.pushFront(4)

	require.Equal(t, 5, q.popBack())
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.pushFront(3)
	q.pushBack(2)

	require.Equal(t, 2, q.popBack())
	require.Equal(t, 3, q.getFront())
	require.Equal(t, 4, q.getRear())

	require.Equal(t, 4, q.popBack())
	require.Equal(t, 3, q.getFront())
	require.Equal(t, 3, q.getRear())

	require.Equal(t, 3, q.popBack())

	require.True(t, q.isEmpty())
}

func TestGetFrontDequeue(t *testing.T) {
	q := dequeue{}

	require.PanicsWithError(t, "dequeue is empty", func() { q.getFront() })

	q.pushFront(4)
	require.Equal(t, 4, q.getFront())

	q.pushBack(5)
	require.Equal(t, 4, q.getFront())

	q.pushFront(7)
	require.Equal(t, 7, q.getFront())

	q.popFront()
	require.Equal(t, 4, q.getFront())

	q.popBack()
	require.Equal(t, 4, q.getFront())
}

func TestGetRearDequeue(t *testing.T) {
	q := dequeue{}

	require.PanicsWithError(t, "dequeue is empty", func() { q.getRear() })

	q.pushBack(4)
	require.Equal(t, 4, q.getRear())

	q.pushBack(5)
	require.Equal(t, 5, q.getRear())

	q.popBack()
	require.Equal(t, 4, q.getRear())

	q.pushFront(10)
	require.Equal(t, 4, q.getRear())

	q.popFront()
	require.Equal(t, 4, q.getRear())
}

func TestIsDequeueEmpty(t *testing.T) {
	q := dequeue{}

	require.True(t, q.isEmpty())

	q.pushBack(4)
	q.pushFront(5)
	require.False(t, q.isEmpty())

	require.Equal(t, 4, q.popBack())
	require.Equal(t, 5, q.popFront())

	require.True(t, q.isEmpty())
}

func TestDequeSize(t *testing.T) {
	q := dequeue{}

	require.Equal(t, 0, q.size())

	q.pushBack(4)
	q.pushFront(5)
	require.Equal(t, 2, q.size())

	q.popBack()
	q.popFront()
	require.Equal(t, 0, q.size())
}
