package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	var st = stack{}

	st.push(2)
	require.Equal(t, 2, st.peek())

	st.push(4)
	require.Equal(t, 4, st.peek())
}

func TestPeek(t *testing.T) {
	var st = stack{}

	require.PanicsWithError(t, "stack is empty", func() { st.peek() })

	st.push(2)
	require.Equal(t, 2, st.peek())

	require.Equal(t, 2, st.peek())
}

func TestPop(t *testing.T) {
	var st = stack{}

	st.push(2)
	st.push(4)
	st.push(8)
	st.push(16)

	require.Equal(t, 16, st.pop())
	require.Equal(t, 8, st.pop())
	require.Equal(t, 4, st.pop())
	require.Equal(t, 2, st.peek())
	require.Equal(t, 2, st.pop())

	require.PanicsWithError(t, "stack is empty", func() { st.pop() })
}

func TestIsEmpty(t *testing.T) {
	var st = stack{}

	require.Equal(t, true, st.isEmpty())

	st.push(2)
	require.Equal(t, false, st.isEmpty())

	st.pop()
	require.Equal(t, true, st.isEmpty())
}

func TestLength(t *testing.T) {
	var st = stack{}

	require.Equal(t, 0, st.length())

	st.push(1)
	require.Equal(t, 1, st.length())

	st.push(4)
	require.Equal(t, 2, st.length())

	st.pop()
	require.Equal(t, 1, st.length())

	st.pop()
	require.Equal(t, 0, st.length())
}
