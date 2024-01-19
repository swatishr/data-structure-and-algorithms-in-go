package linkedlist

import "fmt"

type node struct {
	val  int
	next *node
}

type SinglyLinkedList struct {
	head *node
}

// inserts at the end
func (s *SinglyLinkedList) Insert(newVal int) {
	newNode := &node{
		val:  newVal,
		next: nil,
	}

	if s.head == nil {
		s.head = newNode
		return
	}

	p := s.head
	for p.next != nil {
		p = p.next
	}
	p.next = newNode
}

// inserts at a given position. position starts from 0
func (s *SinglyLinkedList) InsertAtPos(newVal int, pos int) {
	newNode := &node{
		val:  newVal,
		next: nil,
	}

	if pos == 0 {
		newNode.next = s.head
		s.head = newNode
		return
	}

	counter := 1
	p := s.head
	for counter < pos {
		if p == nil {
			fmt.Println("invalid position")
			return
		}
		p = p.next
		counter += 1
	}
	p.next = newNode
}

func (s *SinglyLinkedList) Delete(deleteVal int) {
	p := s.head
	var prev *node
	for p != nil {
		if p.val == deleteVal {
			if p == s.head {
				s.head = p.next
				return
			}
			prev.next = p.next
			return
		}
		prev = p
		p = p.next
	}
	fmt.Printf("%d not present\n", deleteVal)
}

func (s *SinglyLinkedList) Length() int {
	p := s.head
	length := 0

	for p != nil {
		p = p.next
		length += 1
	}

	return length
}

func (s *SinglyLinkedList) Print() {
	p := s.head

	fmt.Printf("\n[")
	for p != nil {
		fmt.Printf("%d, ", p.val)
		p = p.next
	}
	fmt.Printf("]\n")
}

// search the element and return its position
func (s *SinglyLinkedList) RecursiveSearch(searchVal int) int {
	return recursiveSearchHelper(searchVal, s.head, 0)
}

func recursiveSearchHelper(searchVal int, curr *node, pos int) int {
	if curr == nil {
		return -1
	}

	if curr.val == searchVal {
		return pos
	}

	return recursiveSearchHelper(searchVal, curr.next, pos+1)
}

func (s *SinglyLinkedList) IterativeSearch(searchVal int) int {

	if s.head == nil {
		return -1
	}

	p := s.head
	pos := 0

	for p != nil {
		if p.val == searchVal {
			return pos
		}
		p = p.next
		pos += 1
	}

	return -1
}

func (s *SinglyLinkedList) RecursiveReverse() {
	s.head = recursiveReverseHelper(s.head, nil)
}

func recursiveReverseHelper(curr, prev *node) *node {
	if curr == nil {
		return prev
	}

	head := recursiveReverseHelper(curr.next, curr)
	curr.next = prev
	return head
}

func (s *SinglyLinkedList) IterativeReverse() {
	var prev, curr, next *node

	if s.head == nil {
		return
	}

	prev = nil
	curr = s.head
	next = s.head.next

	for next != nil {
		curr.next = prev
		prev = curr
		curr = next
		next = curr.next
	}

	curr.next = prev
	s.head = curr
}
