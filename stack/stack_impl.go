package stack

import "fmt"

type node struct {
	val  int
	next *node
}

type stack struct {
	top *node
	len int
}

func (s *stack) push(x int) {
	curr := &node{
		val:  x,
		next: s.top,
	}
	s.top = curr
	s.len += 1
}

func (s *stack) peek() int {
	if s.isEmpty() {
		panic(fmt.Errorf("stack is empty"))
	}
	return s.top.val
}

func (s *stack) pop() int {
	if s.isEmpty() {
		panic(fmt.Errorf("stack is empty"))
	}

	popped := s.top
	s.top = popped.next

	s.len -= 1

	return popped.val
}

func (s *stack) isEmpty() bool {
	return s.top == nil
}

func (s *stack) length() int {
	return s.len
}
