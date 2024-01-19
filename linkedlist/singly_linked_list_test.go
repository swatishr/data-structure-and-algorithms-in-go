package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	var testSinglyLL = &SinglyLinkedList{}

	require.Equal(t, 0, testSinglyLL.Length())

	testSinglyLL.Insert(10)
	testSinglyLL.Insert(20)
	testSinglyLL.Insert(30)
	testSinglyLL.Insert(40)

	require.Equal(t, 4, testSinglyLL.Length())

	testSinglyLL.Print()
}

func TestInsertAtPos(t *testing.T) {
	var testSinglyLL = &SinglyLinkedList{}

	require.Equal(t, 0, testSinglyLL.Length())

	testSinglyLL.Insert(10)
	testSinglyLL.Insert(20)
	testSinglyLL.InsertAtPos(30, 0)
	require.Equal(t, 30, testSinglyLL.head.val)

	testSinglyLL.InsertAtPos(40, 3)
	testSinglyLL.InsertAtPos(50, 8)

	require.Equal(t, 4, testSinglyLL.Length())

	require.Equal(t, 0, testSinglyLL.IterativeSearch(30))
	require.Equal(t, 3, testSinglyLL.RecursiveSearch(40))
	require.Equal(t, -1, testSinglyLL.IterativeSearch(50))

	testSinglyLL.Print()
}

func TestDelete(t *testing.T) {
	var testSinglyLL = &SinglyLinkedList{}

	require.Equal(t, 0, testSinglyLL.Length())

	testSinglyLL.Insert(10)
	testSinglyLL.Insert(20)
	testSinglyLL.Insert(30)
	require.Equal(t, 3, testSinglyLL.Length())

	testSinglyLL.Delete(20)
	require.Equal(t, 2, testSinglyLL.Length())

	testSinglyLL.Delete(10)
	require.Equal(t, 1, testSinglyLL.Length())
	require.Equal(t, 30, testSinglyLL.head.val)

	testSinglyLL.Delete(50)

	require.Equal(t, -1, testSinglyLL.IterativeSearch(20))
	require.Equal(t, -1, testSinglyLL.RecursiveSearch(10))

	testSinglyLL.Print()
}

func TestReverse(t *testing.T) {
	var testSinglyLL = &SinglyLinkedList{}

	testSinglyLL.Insert(10)
	testSinglyLL.Insert(20)
	testSinglyLL.Insert(30)
	testSinglyLL.Insert(40)

	t.Log("---Before Reversal---")
	testSinglyLL.Print()
	require.Equal(t, 10, testSinglyLL.head.val)

	testSinglyLL.IterativeReverse()
	require.Equal(t, 40, testSinglyLL.head.val)

	t.Log("---After Reversal---")
	testSinglyLL.Print()

	testSinglyLL.RecursiveReverse()
	require.Equal(t, 10, testSinglyLL.head.val)

	t.Log("---Second Reversal - Back to original---")
	testSinglyLL.Print()
}
