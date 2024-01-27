package queues

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPQEnqueue(t *testing.T) {

	q := priorityQueue{}

	require.NotPanics(t, func() { q.enqueue(4, 3) })
	require.Equal(t, 4, q.getFront())
	require.Equal(t, 4, q.getRear())

	// should be front
	q.enqueue(5, 4)
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 4, q.getRear())

	// should be at the end
	q.enqueue(16, 0)
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 16, q.getRear())

	//should be in between
	q.enqueue(10, 2)
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 16, q.getRear())

	// priority equal, so should be fifo
	q.enqueue(13, 4)
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 16, q.getRear())
}

func TestPQDequeue(t *testing.T) {

	q := priorityQueue{}

	require.PanicsWithError(t, "no data to dequeue", func() { q.dequeue() })

	q.enqueue(5, 7)
	q.enqueue(4, 8)

	require.Equal(t, 4, q.dequeue())
	require.Equal(t, 5, q.getFront())
	require.Equal(t, 5, q.getRear())

	q.enqueue(3, 6)
	q.enqueue(2, 7)

	require.Equal(t, 5, q.dequeue())
	require.Equal(t, 2, q.getFront())
	require.Equal(t, 3, q.getRear())

	require.Equal(t, 2, q.dequeue())
	require.Equal(t, 3, q.getFront())
	require.Equal(t, 3, q.getRear())

	require.Equal(t, 3, q.dequeue())

	require.True(t, q.isEmpty())
}

func TestPQGetFront(t *testing.T) {
	q := priorityQueue{}

	require.PanicsWithError(t, "queue is empty", func() { q.getFront() })

	q.enqueue(4, 10)
	require.Equal(t, 4, q.getFront())

	q.enqueue(5, 6)
	require.Equal(t, 4, q.getFront())

	q.dequeue()
	require.Equal(t, 5, q.getFront())
}

func TestPQGetRear(t *testing.T) {
	q := priorityQueue{}

	require.PanicsWithError(t, "queue is empty", func() { q.getRear() })

	q.enqueue(4, 10)
	require.Equal(t, 4, q.getRear())

	q.enqueue(5, 7)
	require.Equal(t, 5, q.getRear())

	q.dequeue()
	require.Equal(t, 5, q.getRear())
}

func TestPQIsEmpty(t *testing.T) {
	q := priorityQueue{}

	require.True(t, q.isEmpty())

	q.enqueue(4, 10)
	q.enqueue(5, 7)
	require.False(t, q.isEmpty())

	q.dequeue()
	q.dequeue()
	require.True(t, q.isEmpty())

}

func TestPQLength(t *testing.T) {
	q := priorityQueue{}

	require.Equal(t, 0, q.length())

	q.enqueue(4, 10)
	q.enqueue(5, 11)
	require.Equal(t, 2, q.length())

	q.dequeue()
	q.dequeue()
	require.Equal(t, 0, q.length())

}
