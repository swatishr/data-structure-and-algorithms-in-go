package queues

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {

	q := queue{}

	require.NotPanics(t, func() { q.enqueue(4) })
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.enqueue(5)
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 5, q.getRear())
}

func TestDequeue(t *testing.T) {

	q := queue{}

	require.PanicsWithError(t, "no data to dequeue", func() { q.dequeue() })

	q.enqueue(5)
	q.enqueue(4)

	require.Equal(t, 5, q.dequeue())
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	q.enqueue(3)
	q.enqueue(2)

	require.Equal(t, 4, q.dequeue())
	require.Equal(t, 3, q.getFront())
	require.Equal(t, 2, q.getRear())

	require.Equal(t, 3, q.dequeue())
	require.Equal(t, 2, q.getFront())
	require.Equal(t, 2, q.getRear())

	require.Equal(t, 2, q.dequeue())

	require.True(t, q.isEmpty())
}

func TestGetFront(t *testing.T) {
	q := queue{}

	require.PanicsWithError(t, "queue is empty", func() { q.getFront() })

	q.enqueue(4)
	require.Equal(t, 4, q.getFront())

	q.enqueue(5)
	require.Equal(t, 4, q.getFront())

	q.dequeue()
	require.Equal(t, 5, q.getFront())
}

func TestGetRear(t *testing.T) {
	q := queue{}

	require.PanicsWithError(t, "queue is empty", func() { q.getRear() })

	q.enqueue(4)
	require.Equal(t, 4, q.getRear())

	q.enqueue(5)
	require.Equal(t, 5, q.getRear())

	q.dequeue()
	require.Equal(t, 5, q.getRear())
}

func TestIsEmpty(t *testing.T) {
	q := queue{}

	require.True(t, q.isEmpty())

	q.enqueue(4)
	q.enqueue(5)
	require.False(t, q.isEmpty())

	q.dequeue()
	q.dequeue()
	require.True(t, q.isEmpty())

}

func TestLength(t *testing.T) {
	q := queue{}

	require.Equal(t, 0, q.length())

	q.enqueue(4)
	q.enqueue(5)
	require.Equal(t, 2, q.length())

	q.dequeue()
	q.dequeue()
	require.Equal(t, 0, q.length())

}
